/* eslint-disable camelcase */
import Vue from 'vue'
import { ValidationProvider, ValidationObserver, localize, extend } from 'vee-validate'
import ja from 'vee-validate/dist/locale/ja.json'
import { required, confirmed, max, min, email, image, alpha_dash, size } from 'vee-validate/dist/rules'
import { RuleParamSchema, ValidationRule } from 'vee-validate/dist/types/types'
import dayjs from '~/plugins/dayjs'

Vue.component('ValidationProvider', ValidationProvider)
Vue.component('ValidationObserver', ValidationObserver)

const hiraganaCustomRule: ValidationRule = {
  params: ['hiragana'],
  validate: (value: string): boolean => {
    const pattern: RegExp = /^[ぁ-ゔー]*$/g
    return pattern.test(value)
  },
  message: (field: string): string => {
    return `${field}はひらがなのみ使用できます`
  },
}

const passwordCustomRule: ValidationRule = {
  params: ['password'],
  validate: (value: string): boolean => {
    const pattern: RegExp = /^[a-zA-Z0-9_!@#$_%^&*.?()-=+]*$/g
    return pattern.test(value)
  },
  message: (field: string): string => {
    return `${field}は英数字と_!@#$_%^&*.?()-=+のみ使用できます`
  },
}

const afterCustomRule: ValidationRule = {
  params: ['target'] as RuleParamSchema[],
  validate: (value: string, params: Record<string, any>) => {
    const target = params['target']
    const date1 = dayjs(value)
    const date2 = dayjs(target)
    if (!date1 || !date2) return false
    return date1.isAfter(date2)
  },
  message: '{_field_}は、{target}よりあとの日付を入力してください',
}

// Basic Rules
extend('required', { ...required })
extend('confirmed', { ...confirmed })
extend('max', { ...max })
extend('min', { ...min })
extend('email', { ...email })
extend('image', { ...image })
extend('alpha_dash', { ...alpha_dash })
extend('size', { ...size })

// Custom Rules
extend('hiragana', { ...hiraganaCustomRule })
extend('password', { ...passwordCustomRule })
extend('after', { ...afterCustomRule })

localize('ja', ja)
