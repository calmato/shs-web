<template>
  <v-container class="mt-4">
    <p class="text-h5 mb-2">開講科目設定</p>
    <div v-for="(schoolType, i) in schoolTypeArray" :key="i" class="mb-8">
      <div class="d-flex align-center mb-2">
        <p class="text-subtitle-1 mr-4 mb-0">開講科目一覧（{{ schoolType }}）</p>
        <v-tooltip right>
          <template #activator="{ on, attrs }">
            <v-btn
              class="suffix-btn"
              color="primary"
              fab
              elevation="0"
              v-bind="attrs"
              v-on="on"
              @click="handleAddButton(schoolType)"
            >
              <v-icon>mdi-plus</v-icon>
            </v-btn>
          </template>
          <span>科目を新規作成</span>
        </v-tooltip>
      </div>
      <v-chip
        v-for="subject in subjectsMap[schoolType]"
        :key="subject.id"
        :color="subject.color"
        class="ma-1"
        @click="handleClickSubjectChip(subject)"
      >
        {{ subject.name }}
      </v-chip>
    </div>

    <v-dialog v-model="isOpen" @click:outside="handleDialoOutsideClick">
      <v-card>
        <v-card-title class="primary white--text">
          {{ `開講科目${dialogType === '新規作成' ? '追加' : '編集'}（${formData.schoolType}）` }}
          <v-spacer />
          <v-btn
            v-if="dialogType === '編集'"
            class="suffix-btn"
            fab
            elevation="0"
            color="primary"
            @click="handleDeleteButton"
          >
            <v-icon>mdi-delete</v-icon>
          </v-btn>
        </v-card-title>
        <v-card-text class="pt-4">
          <p class="text-body-1">
            {{ `${dialogType === '新規作成' ? '追加' : '編集'}する科目の情報を入力してください。` }}
          </p>
          <v-text-field v-model="formData.name" label="科目名"></v-text-field>
          <p class="v-label mb-0">タグの色</p>
          <v-color-picker
            v-model="formData.color"
            hide-sliders
            hide-canvas
            hide-inputs
            show-swatches
            :swatches="swatches"
          />
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn v-if="dialogType === '新規作成'" color="primary" @click="handleSubmitButton">追加</v-btn>
          <v-btn v-if="dialogType === '編集'" color="primary" @click="handleEditSubmitButton">更新</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, ref, computed, useStore, reactive } from '@nuxtjs/composition-api'
import { CommonStore, LessonStore } from '~/store'
import { SubjectNewForm } from '~/types/form'
import { Subject } from '~/types/store'

export default defineComponent({
  setup() {
    const store = useStore()

    const schoolTypeArray: string[] = ['小学校', '中学校', '高校']
    const swatches: string[][] = [['#F8BBD0'], ['#DBD0E6'], ['#BBDEFB'], ['#E8F5E9'], ['#FFE0B2']]

    const isOpen = ref<boolean>(false)
    const dialogType = ref<'新規作成' | '編集'>('新規作成')

    const selectedSubjectId = ref<number>(-1)

    const formData = reactive<SubjectNewForm>({
      name: '',
      schoolType: '小学校',
      color: swatches[0][0],
    })

    const subjects = computed<Subject[]>(() => store.getters['lesson/getSubjects'])
    const subjectsMap = computed<{ [key in '小学校' | '中学校' | '高校']: Subject[] }>(() => {
      return {
        小学校: subjects.value.filter((item) => item.schoolType === '小学校'),
        中学校: subjects.value.filter((item) => item.schoolType === '中学校'),
        高校: subjects.value.filter((item) => item.schoolType === '高校'),
      }
    })

    const handleDialogOutsideClick = () => {
      formData.name = ''
      formData.color = swatches[0][0]
    }

    const handleClickSubjectChip = (subject: Subject) => {
      dialogType.value = '編集'
      isOpen.value = true
      selectedSubjectId.value = subject.id
      formData.name = subject.name
      formData.schoolType = subject.schoolType
      formData.color = subject.color
    }

    const handleAddButton = (schoolType: '小学校' | '中学校' | '高校') => {
      dialogType.value = '新規作成'
      isOpen.value = true
      formData.schoolType = schoolType
    }

    const handleDeleteButton = async () => {
      if (selectedSubjectId.value === -1) {
        return
      }
      try {
        await LessonStore.deleteSubject(selectedSubjectId.value)
        selectedSubjectId.value = -1
        isOpen.value = false
        CommonStore.showSuccessInSnackbar(`「${formData.name}」を削除しました。`)
      } catch (err) {
        CommonStore.showErrorInSnackbar(err)
      }
    }

    const handleEditSubmitButton = async () => {
      if (selectedSubjectId.value === -1) {
        return
      }
      try {
        await LessonStore.editSubject({ ...formData, subjectId: selectedSubjectId.value })
        selectedSubjectId.value = -1
        isOpen.value = false
        CommonStore.showSuccessInSnackbar(`「${formData.name}」を更新しました。`)
      } catch (err) {
        CommonStore.showErrorInSnackbar(err)
      }
    }

    const handleSubmitButton = async () => {
      try {
        await LessonStore.createSubject(formData)
        CommonStore.showSuccessInSnackbar(`「${formData.name}」を作成しました。`)
        isOpen.value = false
        formData.name = ''
        formData.color = swatches[0][0]
      } catch (err) {
        CommonStore.showErrorInSnackbar(err)
      }
    }

    return {
      schoolTypeArray,
      swatches,
      isOpen,
      dialogType,
      formData,
      subjectsMap,
      handleDialogOutsideClick,
      handleClickSubjectChip,
      handleDeleteButton,
      handleAddButton,
      handleSubmitButton,
      handleEditSubmitButton,
    }
  },
})
</script>

<style lang="scss" scoped>
.suffix-btn {
  height: 24px;
  width: 24px;
}

.v-color-picker {
  ::v-deep > .v-color-picker__swatches > div {
    justify-content: flex-start;
  }
}
</style>
