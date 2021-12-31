<template>
  <validation-provider v-slot="{ errors, valid }" :name="label" :vid="name" :rules="rules">
    <v-select
      v-model="formData"
      :items="items"
      :label="label"
      :error-messages="errors"
      :success="valid"
      :autofocus="autofocus"
      :outlined="outlined"
      :chips="chips"
      :append-outer-icon="appendOuterIcon"
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
    chips: {
      type: Boolean,
      required: false,
      default: false,
    },
    rules: {
      type: Object,
      required: false,
      default: () => ({}),
    },
    items: {
      type: Array,
      required: false,
      default: () => [],
    },
    value: {
      type: [String, Number],
      required: false,
      default: undefined,
      type: [String, Number],
    },
    appendOuterIcon: {
      required: false,
      default: undefined,
      type: String,
    },
  },

  setup(props, { emit }: SetupContext) {
    const formData = computed({
      get: (): any => props.value,
      set: (val: any) => emit('update:value', val),
    })

    return {
      formData,
    }
  },
})
</script>
