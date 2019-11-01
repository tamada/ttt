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
