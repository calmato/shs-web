<template>
  <div class="pa-2">
    <p class="text-h5 mb-0">メールアドレス変更</p>
    <the-form-group class="pa-2">
      <the-text-field
        :label="form.options.mail.label"
        :rules="form.options.mail.rules"
        :value.sync="form.params.updateMail"
        type="email"
      />
    </the-form-group>
    <div class="d-flex justify-end pr-4">
      <v-btn color="primary"> 保存する </v-btn>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from '@vue/composition-api'
import TheFormGroup from '~/components/atoms/TheFormGroup.vue'
import TheTextField from '~/components/atoms/TheTextField.vue'
import { TeacherUpdateMailForm, TeacherUpdateMailOptions, TeacherUpdateMailParams } from '~/types/form'

export default defineComponent({
  components: {
    TheFormGroup,
    TheTextField,
  },

  props: {
    form: {
      type: Object as PropType<TeacherUpdateMailForm>,
      default: () => ({
        params: TeacherUpdateMailParams,
        options: TeacherUpdateMailOptions,
      }),
    },
  },

  setup() {},
})
</script>
