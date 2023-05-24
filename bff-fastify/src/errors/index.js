class CustomError extends Error {
    constructor({
                    statusCode = 500,
                    message = '',
                    errorCode = 10000
                }) {
        super()
        this.statusCode = statusCode
        this.message = `${message}`
        this.errorCode = errorCode
    }

    async send(req, res, correlationId){
        const response =
            JSON.stringify({
                statusCode: this.statusCode,
                message: this.message,
                errorCode: this.errorCode,
                correlationId
            })
        console.log(response)

        return res.status(this.statusCode).send(response)
    }
}

module.exports = CustomError
