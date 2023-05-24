const client = require('../../utils/httpHelper')
const config = require('../../config')
const CustomError = require('../../errors')
const { isEmptyObject } = require('../../utils')

exports.getAllCustomers = async ({ limit, offset, orderDirection, orderBy, isCount }) => {
    const query = { limit, offset, orderDirection, orderBy, isCount }
    // const url = '/customers?' + Object.keys(query).filter(k => !!query[k]).map(k => `${k}=${encodeURIComponent(query[k])}`).join('&')
    const url = '/customers?' + Object.keys(query).filter(k => !!query[k]).map(k => `${k}=${query[k]}`).join('&')

    const response = await client.get({
        baseURL : config.services.internals.customerApi,
        url
        })

    if (isEmptyObject(response) || response?.status >= 500) {
        throw new CustomError({
            statusCode: 500,
            message: 'Customer service failed to get all customers!',
            errorCode: 10101
        })
    } else if (response?.status === 400) {
        throw new CustomError({
            statusCode: 400,
            message: response?.data.Desc,
            errorCode: 10102
        })
    }

    return response?.data
}
