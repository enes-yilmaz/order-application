const clients = require('../../clients')

exports.getAllCustomers = async ({ limit, offset, orderDirection, orderBy, isCount }) => {
    return await clients.customerApi.getAllCustomers({ limit, offset, orderDirection, orderBy, isCount })
}

exports.getCustomer = async (customerId) => {
    return await clients.customerApi.getCustomer(customerId)
}

exports.getCustomerValidation = async (customerId) => {
    return await clients.customerApi.getCustomerValidation(customerId)
}

exports.postCreateCustomer = async (body) => {
    return await clients.customerApi.postCreateCustomer(body)
}

exports.postUpdateCustomer = async (body) => {
    return await clients.customerApi.postUpdateCustomer(body)
}

exports.deleteCustomer = async (customerId) => {
    return await clients.customerApi.deleteCustomer(customerId)
}
