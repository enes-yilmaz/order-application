const services = require('../../services/index')
const helpers = require('./helpers')

exports.getAllCustomers = async (req, res) => {
    let {
        limit,
        offset,
        orderDirection,
        orderBy,
        isCount
    } = req.query

    limit = limit < 100 && limit > 0 ? +limit : 25
    offset = offset >= 0 ? +offset : 0

    const resp = await services.customerApi.getAllCustomers({ limit, offset, orderDirection, orderBy, isCount })
    return res.status(200).send({ ...resp })
}
