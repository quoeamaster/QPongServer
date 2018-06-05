/**
 *  PS. a service or util to handle disk I/O operation. However, now is just
 *  a package of I/O functions...
 */
package util

import (
    "mime/multipart"
    "strings"
    "fmt"
    "bytes"
    "io/ioutil"
)

/**
 *  struct to wrap up a bunch of meta data about the Storage
 */
type StorageMeta struct {
    FileOrPath string
}
/**
 *  create a new StorageMeta instance
 */
func NewStorageMeta(fileOrPath string) StorageMeta {
    meta := StorageMeta{}
    meta.FileOrPath = fileOrPath
    return meta
}

func IsStorageMetaValid(meta *StorageMeta) bool {
    if meta == nil {
        return false
    }
    if strings.Compare("", strings.TrimSpace(meta.FileOrPath)) ==0 {
        return false
    }
    // More rules to come if the StorageMeta is more complex later on
    return true
}


/**
 *  write the given multipart file to disc.
 *  StorageMeta wraps the details on where the file contents should be
 *  persisted to.
 */
func WriteMultiPartFileToDisc(file *multipart.File, meta StorageMeta) (bytesWrote int, err error) {
    var wBuffer bytes.Buffer
    iBytesWrote := 0

    if !IsStorageMetaValid(&meta) {
        err = fmt.Errorf("the supplied StorageMeta is not valid; please check [%v]", meta)
        return -1, err
    }

    if file != nil {
        // 1 kb of bytes...
        bBytes := make([]byte, 10000)

        fileVal := *file
        for true {
            iBytes, err := fileVal.Read(bBytes)
            // break out as already read all the data from the file
            if iBytes == 0 {
                break
            }
            if err != nil {
                return iBytes, err
            }

            // add the []byte to a buffer for later write operation
            iBytes, err = wBuffer.Write(bBytes[0: iBytes])
            if err != nil {
                return iBytes, err
            }
            // update the value of bytes wrote to the buffer
            iBytesWrote += iBytes
        }   // end -- for

        // final write to the destination noted by the StorageMeta
        err = ioutil.WriteFile(meta.FileOrPath, wBuffer.Bytes(), 0644)
        if err != nil {
            return -1, err
        }
        // everything is fine
        bytesWrote = iBytesWrote
        err = nil

    } else {
        err = fmt.Errorf("the multipart.file is invalid; please check [%v]", *file)
        return -1, err
    }
    return bytesWrote, err
}



