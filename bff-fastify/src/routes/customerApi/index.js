const customerController = require('../../controllers/customerApi/index')
const schema = require('./schema')

const customerRoutes = [
    {
        method: 'GET',
        url: '/api/v1/customers',
        handler: customerController.getAllCustomers,
        schema: schema.getAllCustomers
    }
]

module.exports = customerRoutes
