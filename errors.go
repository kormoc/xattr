package xattr

func syscallErrorToXAttrError(err error) error {
    if err == nil {
        return nil
    }
    switch err.Error() {
        case "attribute not found":
            return XAttrErrorAttributeNotFound{}
        case "errno 0":
            return nil
        case "no data available":
            return XAttrErrorNoDataAvailable{}
        case "result too large":
            return XAttrErrorResultTooLarge{}
        default:
            return err
    }
}

type XAttrErrorAttributeNotFound struct { error }

func (e XAttrErrorAttributeNotFound) Error() string {
    return "attribute not found"
}

type XAttrErrorNoDataAvailable struct { error }

func (e XAttrErrorNoDataAvailable) Error() string {
    return "no data available"
}

type XAttrErrorResultTooLarge struct { error }

func (e XAttrErrorResultTooLarge) Error() string {
    return "result too large"
}

func XAttrErrorIsFatal(err error) bool {
    switch err.(type) {
        case nil:
            return false
        case XAttrErrorAttributeNotFound:
            return false
        case XAttrErrorNoDataAvailable:
            return false
        case XAttrErrorResultTooLarge:
            return true
        default:
            return true
    }
}
