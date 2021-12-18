<template>
  <v-container fill-height>
    <v-row class="justify-center align-center">
      <v-col cols="12" sm="8">
        <v-card>
          <v-toolbar color="primary" dark flat>SHS Web 講師ログイン画面</v-toolbar>
          <v-card-text class="my-4">
            <v-text-field v-model="signInForm.mail" label="メールアドレス" type="email" autofocus outlined />
            <v-text-field v-model="signInForm.password" label="パスワード" type="password" outlined />
            <v-btn color="primary" block :loading="loading" :disabled="loading" @click="onClick">サインイン</v-btn>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { computed, defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import { SignInForm } from '~/types/form'

export default defineComponent({
  props: {
    form: {
      type: Object as PropType<SignInForm>,
      default: () => ({
        mail: '',
        password: '',
      }),
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

    const onClick = () => {
      emit('click')
    }

    return {
      signInForm,
      onClick,
    }
  },
})
</script>
