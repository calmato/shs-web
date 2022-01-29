<template>
  <div class="pa-2">
    <p class="text-h5 mb-0">メールアドレス変更</p>
    <the-form-group class="pa-2">
      <the-text-field
        :label="updateMailForm.options.mail.label"
        :rules="updateMailForm.options.mail.rules"
        :value.sync="updateMailForm.params.mail"
        type="email"
      />
    </the-form-group>
    <div class="d-flex justify-end pr-4">
      <v-btn color="primary" @click="handleSubmit"> 保存する </v-btn>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, SetupContext } from '@vue/composition-api'
import TheFormGroup from '~/components/atoms/TheFormGroup.vue'
import TheTextField from '~/components/atoms/TheTextField.vue'
import { CommonStore, UserStore } from '~/store'
import { TeacherUpdateMailOptions, TeacherUpdateMailParams } from '~/types/form'

export default defineComponent({
  components: {
    TheFormGroup,
    TheTextField,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router
    const updateMailForm = reactive({
      params: TeacherUpdateMailParams,
      options: TeacherUpdateMailOptions,
    })

    const handleSubmit = async (): Promise<void> => {
      CommonStore.startConnection()

      await UserStore.updateMail({ form: updateMailForm })
        .then(() => {
          router.push('/settings')
          CommonStore.showSnackbar({ color: 'success', message: 'メールアドレスを更新しました。' })
        })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    return {
      updateMailForm,
      handleSubmit,
    }
  },
})
</script>
