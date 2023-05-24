const customerRoutes = require('./customerApi/index')
const orderRoutes = require('./orderApi/index')

const routes = [
    ...customerRoutes,
    ...orderRoutes
]

module.exports = routes
