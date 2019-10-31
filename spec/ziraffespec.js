'use strict'

const Ziraffe = require('../ziraffe')
const ziraffe = new Ziraffe()

describe('似た講義名を探す', () => {
    it('統計解析に似た講義名を探す', () => {
        const array = ziraffe.similarLectures('電気回路')
        expect(1).toBe(array.length)
    })
})

describe('編集距離を計算する', () => {
    it('distanceとsimilarity', () => {
        expect(8).toBe(ziraffe.calculateEditDistance('distance', 'similarity').distance)
    })
})
