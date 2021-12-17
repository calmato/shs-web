<template>
  <!-- TODO: デザインの整形 -->
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8">
      <v-text-field v-model="signInForm.mail" label="メールアドレス" autofocus outlined />
      <v-text-field v-model="signInForm.password" label="パスワード" outlined />
      <v-btn color="primary" block :loading="loading" :disabled="loading" @click="onClick">サインイン</v-btn>
    </v-col>
  </v-row>
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
