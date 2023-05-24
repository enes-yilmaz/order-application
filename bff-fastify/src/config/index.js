const qaConfig = require('./qa-config')
const prodConfig = require('./prod-config')

const config = {
    ...qaConfig,
    ...prodConfig
}

const chosenEnv = process.env.ENV || 'qa'

console.log(`bff-fastify runs on ${chosenEnv}`)

module.exports = config[chosenEnv]
