const CustomError = require('../../errors/index')

exports.validateDates = ({ from, to }) => {
    const now = new Date()
    now.setHours(0, 0, 0, 0)
    const nDaysBefore = new Date(now)
    nDaysBefore.setDate(nDaysBefore.getDate() - 60)

    if (to) {
        to = new Date(to)
        if (to.toString() === 'Invalid Date') {
            throw new CustomError({
                statusCode: 400,
                message: 'Failed parsing endDate',
                errorCode: 70007
            })
        }
    } else {
        to = new Date()
    }

    to.setHours(0, 0, 0, 0)

    if (from) {
        from = new Date(from)
        if (from.toString() === 'Invalid Date') {
            throw new CustomError({
                statusCode: 400,
                message: 'Failed parsing startDate',
                errorCode: 70008
            })
        }
    } else {
        from = new Date()
        from.setDate(from.getDate() - 60)
    }

    from.setHours(0, 0, 0, 0)

    if (from.getTime() > to.getTime()) {
        throw new CustomError({
            statusCode: 400,
            message: 'from parameter can\'t be bigger than to parameter',
            errorCode: 70013
        })
    } else if (from.getTime() < nDaysBefore.getTime()) {
        throw new CustomError({
            statusCode: 400,
            message: 'from parameter must be greater than 60 days from now',
            errorCode: 70014
        })
    } else if (to.getTime() > now.getTime()) {
        throw new CustomError({
            statusCode: 400,
            message: 'to parameter must be less than now',
            errorCode: 70017
        })
    }

    return {
        startDate: from.toISOString(),
        endDate: to.toISOString()
    }
}

exports.getDates = () => {
    const to = new Date()
    to.setHours(0, 0, 0, 0)

    const from = new Date()
    from.setDate(from.getDate() - 70)
    from.setHours(0, 0, 0, 0)

    return {
        startDate: from.toISOString(),
        endDate: to.toISOString()
    }
}
