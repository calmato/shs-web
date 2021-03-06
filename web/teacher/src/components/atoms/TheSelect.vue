<template>
  <validation-provider v-slot="{ errors, valid }" :name="label" :vid="name" :rules="rules">
    <v-select
      v-model="formData"
      :items="items"
      :item-text="itemText"
      :item-value="itemValue"
      :label="label"
      :error-messages="errors"
      :success="isSuccess(valid)"
      :autofocus="autofocus"
      :disabled="disabled"
      :outlined="outlined"
      :multiple="multiple"
      :chips="chips"
      :append-outer-icon="appendOuterIcon"
      @blur="onBlur"
    >
      <slot />
    </v-select>
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
    disabled: {
      type: Boolean,
      required: false,
      default: false,
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
    multiple: {
      type: Boolean,
      required: false,
      default: false,
    },
    rules: {
      type: Object,
      required: false,
      default: undefined,
    },
    items: {
      type: Array,
      required: false,
      default: () => [],
    },
    itemText: {
      type: String,
      required: false,
      default: undefined,
    },
    itemValue: {
      type: String,
      required: false,
      default: undefined,
    },
    value: {
      required: false,
      default: undefined,
      type: [String, Number, Array],
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

    const isSuccess = (valid: boolean): boolean => {
      return props.rules ? valid : false
    }

    const onBlur = (): void => {
      emit('blur')
    }

    return {
      formData,
      isSuccess,
      onBlur,
    }
  },
})
</script>
