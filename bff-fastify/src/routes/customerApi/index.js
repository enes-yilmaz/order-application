const customerController = require('../../controllers/customerApi/index')
const schema = require('./schema')

const customerRoutes = [
    {
        method: 'GET',
        url: '/api/v1/customers',
        handler: customerController.getAllCustomers,
        schema: schema.getAllCustomers
    },
    {
        method: 'GET',
        url: '/api/v1/customers/:customerId',
        handler: customerController.getCustomer,
        schema: schema.getCustomer
    },
    {
        method: 'GET',
        url: '/api/v1/customers/validation/:customerId',
        handler: customerController.getCustomerValidation,
        schema: schema.getCustomerValidation
    },
    {
        method: 'POST',
        url: '/api/v1/customers/create',
        handler: customerController.postCreateCustomer,
        schema: schema.postCreateCustomer
    },
    {
        method: 'POST',
        url: '/api/v1/customers/update',
        handler: customerController.postUpdateCustomer,
        schema: schema.postUpdateCustomer
    },
    {
        method: 'DELETE',
        url: '/api/v1/customers/:customerId',
        handler: customerController.deleteCustomer,
        schema: schema.deleteCustomer
    }
]

module.exports = customerRoutes
