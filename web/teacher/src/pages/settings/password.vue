<template>
  <v-container class="mt-4">
    <div class="pa-2">
      <p class="text-h5 mb-0">パスワードの変更</p>
      <the-form-group class="pa-2">
        <the-text-field
          :label="updatePasswordForm.options.password.label"
          :rules="updatePasswordForm.options.passwordConfirmation.rules"
          :value.sync="updatePasswordForm.params.password"
          type="password"
        />
        <the-text-field
          :label="updatePasswordForm.options.passwordConfirmation.label"
          :rules="updatePasswordForm.options.passwordConfirmation.rules"
          :value.sync="updatePasswordForm.params.passwordConfirmaion"
          type="password"
        />
      </the-form-group>
      <div class="d-flex justify-end pr-4">
        <v-btn color="primary" @click="handleSubmit"> 保存 </v-btn>
      </div>
    </div>
  </v-container>
</template>

<script lang="ts">
import { useRouter } from '@nuxtjs/composition-api'
import { defineComponent, reactive } from '@vue/composition-api'
import TheFormGroup from '~/components/atoms/TheFormGroup.vue'
import TheTextField from '~/components/atoms/TheTextField.vue'
import { CommonStore, UserStore } from '~/store'
import { TeacherUpdatePasswordOptions, TeacherUpdatePasswordParams } from '~/types/form'

export default defineComponent({
  components: {
    TheFormGroup,
    TheTextField,
  },

  setup() {
    const router = useRouter()
    const updatePasswordForm = reactive({
      params: TeacherUpdatePasswordParams,
      options: TeacherUpdatePasswordOptions,
    })

    const handleSubmit = async (): Promise<void> => {
      CommonStore.startConnection()

      await UserStore.updatePassword({ form: updatePasswordForm })
        .then(() => {
          router.push('/signin')
          CommonStore.showSnackbar({ color: 'success', message: 'パスワードを更新しました。' })
        })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    return {
      updatePasswordForm,
      handleSubmit,
    }
  },
})
</script>
