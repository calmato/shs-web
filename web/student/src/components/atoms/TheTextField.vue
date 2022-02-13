<template>
  <validation-provider v-slot="{ errors, valid }" :name="label" :vid="name" :rules="rules">
    <v-text-field
      v-model="formData"
      :type="type"
      :label="label"
      :error-messages="errors"
      :success="isSuccess(valid)"
      :autofocus="autofocus"
      :outlined="outlined"
      :readonly="readonly"
      :height="height"
      :prepend-icon="prependIcon"
    />
  </validation-provider>
</template>

<script lang="ts">
import { defineComponent, computed, SetupContext } from '@nuxtjs/composition-api'

export default defineComponent({
  props: {
    autofocus: {
      type: Boolean,
      required: false,
      default: false,
    },
    height: {
      type: String,
      default: undefined,
    },
    label: {
      type: String,
      required: false,
      default: '',
    },
    name: {
      type: String,
      required: false,
      default: '',
    },
    outlined: {
      type: Boolean,
      required: false,
      default: false,
    },
    prependIcon: {
      type: String,
      required: false,
      default: undefined,
    },
    readonly: {
      type: Boolean,
      require: false,
      default: false,
    },
    rules: {
      type: Object,
      required: false,
      default: undefined,
    },
    type: {
      type: String,
      required: false,
      default: 'text',
    },
    value: {
      type: String,
      required: false,
      default: '',
    },
  },

  setup(props, { emit }: SetupContext) {
    const formData = computed({
      get: () => props.value,
      set: (val: string) => emit('update:value', val),
    })

    const isSuccess = (valid: boolean): boolean => {
      return props.rules ? valid : false
    }

    return {
      formData,
      isSuccess,
    }
  },
})
</script>
