import { setup, refresh } from '~~/test/helpers/store-helper'
import { CommonStore } from '~/store'
import { PromiseState } from '~/types/store'

describe('store/common', () => {
  beforeEach(() => {
    setup()
  })

  afterEach(() => {
    refresh()
  })

  describe('getters', () => {
    it('getPromiseState', () => {
      expect(CommonStore.getPromiseState).toBe(PromiseState.NONE)
    })
  })

  describe('actions', () => {
    describe('startConnection', () => {
      it('stateが更新されること', () => {
        CommonStore.startConnection()
        expect(CommonStore.getPromiseState).toBe(PromiseState.LOADING)
      })
    })

    describe('endConnection', () => {
      it('stateが更新されること', () => {
        CommonStore.endConnection()
        expect(CommonStore.getPromiseState).toBe(PromiseState.NONE)
      })
    })
  })
})
