<template>
  <v-container class="mt-4">
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
import { AuthStore, CommonStore } from '~/store'
import { StudentUpdateMailOptions, StudentUpdateMailParams } from '~/types/form'
export default defineComponent({
  components: {
    TheFormGroup,
    TheTextField,
  },

  setup() {
    const router = useRouter()
    const updateMailForm = reactive({
      params: StudentUpdateMailParams,
      options: StudentUpdateMailOptions,
    })

    const handleSubmit = async (): Promise<void> => {
      CommonStore.startConnection()

      await AuthStore.updateMail({ form: updateMailForm })
        .then(() => {
          AuthStore.signOut()
          router.push('/signin')
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
