<template>
  <v-container>
    <v-row class="justify-center">
      <v-col>
        <div>
          <v-card-title>講師登録</v-card-title>
          <v-card-text>
            <the-teacher-new-form :form="form" />
          </v-card-text>
          <v-card-actions>
            <v-spacer />
            <v-btn color="primary" :loading="loading" :disabled="loading" @click="onSubmit">登録</v-btn>
          </v-card-actions>
        </div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import TheTeacherNewForm from '~/components/organisms/TheTeacherNewForm.vue'
import { TeacherNewForm, TeacherNewOptions, TeacherNewParams } from '~/types/form'

export default defineComponent({
  components: {
    TheTeacherNewForm,
  },

  props: {
    form: {
      type: Object as PropType<TeacherNewForm>,
      default: () => ({
        params: TeacherNewParams,
        options: TeacherNewOptions,
      }),
    },
    loading: {
      type: Boolean,
      default: false,
    },
  },

  setup(_, { emit }: SetupContext) {
    const onSubmit = (): void => {
      emit('submit')
    }

    const onCancel = (): void => {
      emit('cancel')
    }

    return {
      onSubmit,
      onCancel,
    }
  },
})
</script>
