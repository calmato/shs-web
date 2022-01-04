<template>
  <v-menu
    ref="menu"
    :close-on-content-click="true"
    transition="scale-transition"
    offset-y
    max-width="290px"
    min-width="auto"
  >
    <template #activator="{ on, attrs }">
      <validation-provider v-slot="{ errors, valid }" :name="label" :vid="name" :rules="rules">
        <div class="d-flex">
          <v-text-field
            v-model="formData"
            :error-messages="errors"
            :success="valid"
            prepend-icon="mdi-calendar"
            readonly
            v-bind="attrs"
            v-on="on"
          />
          <slot />
        </div>
      </validation-provider>
    </template>
    <v-date-picker v-model="formData" :type="type" :locale="locale" no-title scrollable />
  </v-menu>
</template>

<script lang="ts">
import { computed, defineComponent, SetupContext } from '@nuxtjs/composition-api'

export default defineComponent({
  props: {
    label: {
      type: String,
      required: false,
      default: '',
    },
    locale: {
      type: String,
      require: false,
      default: 'ja',
    },
    name: {
      type: String,
      required: false,
      default: '',
    },
    rules: {
      type: Object,
      required: false,
      default: () => ({}),
    },
    type: {
      type: String,
      require: false,
      default: 'date',
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

    return {
      formData,
    }
  },
})
</script>
