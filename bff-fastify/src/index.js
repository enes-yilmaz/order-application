const config = require('./config/index')
const fastify = require('fastify')({ logger: true, disableRequestLogging: true, ajv: { customOptions: { coerceTypes: 'array' } }, maxParamLength: 900 })
const routes = require('./routes/index')
const swagger = require('./config/swagger')
const CustomError = require('./errors')

const main = async () => {
    await fastify.register(require('@fastify/cors'), {})
    await fastify.register(require('@fastify/swagger'), swagger.options)

    routes.forEach((route, _index) => {
        fastify.route(route)
    })

    fastify.setErrorHandler(function (error, req, res) {
        // Log error
        // Send error response
        if (error?.statusCode) {
            fastify.log.error(error)
            res.status(error?.statusCode).send(error)
        } else {
            const e = new CustomError({
                statusCode: 500,
                message: `An unknown internal error occurred, detail: ${error?.message}`,
                errorCode: error?.errorCode
            })

            fastify.log.error(e)
            res.status(e.statusCode).send(e)

            process.on('unhandledRejection', () => {
                fastify.log.error(e)
                res.status(e.statusCode).send(e)
            })
        }
    })

    try {
        await fastify.ready()
        fastify.swagger()
    } catch (err) {
        fastify.log.error(err)
        process.exit(1)
    }

    const port = 3000
    fastify.listen({ port, host: '0.0.0.0' }, async () => {
        fastify.log.info(`server listening on ${port}`)
    })
}

main()
