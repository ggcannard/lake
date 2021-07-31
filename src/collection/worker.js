#!/usr/bin/env node

require('module-alias/register')
const axios = require('axios')
const _has = require('lodash/has')

const dbConnector = require('@mongo/connection')
const { buildPluginRegistry } = require('../plugins')
const consumer = require('../queue/consumer')
const enrichmentApiUrl = require('@config/resolveConfig').enrichment.connectionString

const queue = 'collection'

const jobHandler = async (job) => {
  const collection = await buildPluginRegistry('collector')
  const {
    db, client
  } = await dbConnector.connect()

  const enrichmentJob = {}

  try {
    if (_has(job, 'jira')) {
      enrichmentJob.jira = await collection.plugins.jiraCollector(db, job.jira)
    }
    if (_has(job, 'gitlab')) {
      enrichmentJob.gitlab = await collection.plugins.gitlabCollector(db, job.gitlab)
    }
  } catch (error) {
    console.log('Failed to collect', error)
  } finally {
    dbConnector.disconnect(client)
  }

  await axios.post(enrichmentApiUrl, enrichmentJob)
}

consumer(queue, jobHandler)
