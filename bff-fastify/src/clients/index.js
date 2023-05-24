const customerApi = require('./customerApi/index')
const orderApi = require('./orderApi/index')

const clients = {
    customerApi,
    orderApi
}

module.exports = clients
