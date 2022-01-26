<template>
  <div class="pa-2">
    <p class="text-h5 mb-0">パスワードの変更</p>
    <the-form-group class="pa-2">
      <the-text-field
        :label="form.options.updatePassword.label"
        :rules="form.options.updatePassword.rules"
        :value.sync="form.params.updatePassword"
        type="password"
      />
      <the-text-field
        :label="form.options.updatePasswordConfirmation.label"
        :rules="form.options.updatePasswordConfirmation.rules"
        :value.sync="form.params.updatePasswordConfirmaion"
        type="password"
      />
    </the-form-group>
    <div class="d-flex justify-end pr-4">
      <v-btn color="primary"> 保存する </v-btn>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive } from '@vue/composition-api'
import TheFormGroup from '~/components/atoms/TheFormGroup.vue'
import TheTextField from '~/components/atoms/TheTextField.vue'
import { TeacherUpdatePasswordOptions, TeacherUpdatePasswordParams } from '~/types/form'

export default defineComponent({
  components: {
    TheFormGroup,
    TheTextField,
  },

  setup() {
    const form = reactive({
      params: TeacherUpdatePasswordParams,
      options: TeacherUpdatePasswordOptions,
    })

    return {
      form,
    }
  },
})
</script>
