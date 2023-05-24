const config = {
    production: {
        env: 'production',
        services: {
            internals: {
                orderApi: 'http://localhost:4001/order-api',
                customerApi: 'http://localhost:4000/customer-api',
            }
        }
    }
}

module.exports = config
