exports.isEmptyObject = (obj = {}) => {
    return !(Object.values(obj).some(x => x))
}
