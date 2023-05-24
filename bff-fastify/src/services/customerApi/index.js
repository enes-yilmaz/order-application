const clients = require('../../clients')

exports.getAllCustomers = async ({ limit, offset, orderDirection, orderBy, isCount }) => {
    return await clients.customerApi.getAllCustomers({ limit, offset, orderDirection, orderBy, isCount })
}
