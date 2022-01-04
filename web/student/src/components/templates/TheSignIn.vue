<template>
  <v-container fill-height>
    <v-row class="justify-center align-center">
      <v-col cols="12" sm="8">
        <the-alert :show="hasError" type="error" @update:show="$emit('update:has-error', $event)">
          メールアドレス もしくは パスワード が間違っています
        </the-alert>
        <v-card>
          <v-form @submit.prevent="onSubmit">
            <v-toolbar color="primary" dark flat>SHS Web 生徒ログイン画面</v-toolbar>
            <v-card-text class="my-4">
              <v-text-field v-model="signInForm.mail" label="メールアドレス" type="email" autofocus outlined />
              <v-text-field v-model="signInForm.password" label="パスワード" type="password" outlined />
              <v-btn color="primary" type="submit" block :loading="loading" :disabled="loading">サインイン</v-btn>
            </v-card-text>
          </v-form>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { computed, defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import TheAlert from '~/components/atoms/TheAlert.vue'
import { SignInForm } from '~/types/form'

export default defineComponent({
  components: {
    TheAlert,
  },

  props: {
    form: {
      type: Object as PropType<SignInForm>,
      default: () => ({
        mail: '',
        password: '',
      }),
    },
    hasError: {
      type: Boolean,
      default: false,
    },
    loading: {
      type: Boolean,
      default: false,
    },
  },

  setup(props, { emit }: SetupContext) {
    const signInForm = computed<SignInForm>({
      get: () => props.form,
      set: (val: SignInForm) => emit('update:form', val),
    })

    const onSubmit = () => {
      emit('click')
    }

    return {
      signInForm,
      onSubmit,
    }
  },
})
</script>
