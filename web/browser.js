'use strict'

const KEY_OF_LOCAL_STORAGE = 'checksOfGotCredits'

const loadChecksFromLocalStorage = () => {
    const gotCredits = localStorage.getItem(KEY_OF_LOCAL_STORAGE)
    if (gotCredits != null) {
        gotCredits.split(",").forEach(item => {
            const elements = document.getElementsByTagName('input')
            Array.prototype.filter.call(elements, inputItem => inputItem.value === item)
                .forEach(inputItem => inputItem.checked = true)
        })
    }
}

const downloadCreditsJSON = () => {
    const array = []
    const elements = document.getElementsByTagName('input')
    Array.prototype.filter.call(elements, inputItem => inputItem.checked)
        .forEach(inputItem => array.push(inputItem.value))
    // refer https://kuroeveryday.blogspot.com/2016/04/byte-order-mark.html
    // Blob表示時に文字化けするのを解決．
    var bom = new Uint8Array([0xEF, 0xBB, 0xBF]);
    // refer https://qiita.com/wadahiro/items/eb50ac6bbe2e18cf8813
    // ファイルを作らずに強制的にダウンロード．
    var blob = new Blob([bom, JSON.stringify(array)], { "type": "text/plain" })
    if (window.navigator.msSaveBlob) {
        window.navigator.msSaveBlob(blob, "credits.json")
    } else {
        window.location.href = window.URL.createObjectURL(blob)
    }
}

const checkAllItems = () => {
    const elements = document.getElementsByTagName('input')
    const array = []
    Array.prototype.forEach.call(elements, inputItem => {
        inputItem.checked = true
        array.push(inputItem.value)
    })
    localStorage.setItem(KEY_OF_LOCAL_STORAGE, array)
}

const clearChecks = () => {
    const elements = document.getElementsByTagName('input')
    Array.prototype.forEach.call(elements, inputItem => {
        inputItem.checked = false
    })
    localStorage.removeItem(KEY_OF_LOCAL_STORAGE)
}

const findGotCredits = () => {
    const elements = document.getElementsByTagName('input')
    const gotCredits = []
    Array.prototype.forEach.call(elements, inputItem => {
        if (inputItem.checked) {
            gotCredits.push(inputItem.value)
        }
    })
    return gotCredits
}

const storeChecksToLocalStorage = (credits) => {
    localStorage.setItem(KEY_OF_LOCAL_STORAGE, credits)
}

const verifyDiploma = () => {
    const credits = findGotCredits()
    storeChecksToLocalStorage(credits)
    checkDiplomaOfCourses(credits)
}

const updateCheckbox = () => {
    const credits = findGotCredits()
    storeChecksToLocalStorage(credits)
}

const initializeEventListener = () => {
    const elements = document.getElementsByTagName("input")
    Array.prototype.forEach.call(elements, item => {
        if (item.type == "checkbox") {
            item.addEventListener('change', updateCheckbox)
        }
    })
}

const initialize = () => {
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)
        .then((result) => {
            go.run(result.instance);
            initDataStore() // call go function via wasm.
            buildHTML()
            loadChecksFromLocalStorage()
            initializeEventListener()
        })
}
