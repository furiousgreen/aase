<template>
  <v-layout justify-start align-start>
    <v-flex xs12>
      <v-card>
        <v-card-title class="headline primary white--text">
          APK File
        </v-card-title>
        <v-card-text class="pa-5">
          <v-form>
            <upload-btn
              block
              fixed-width="100%"
              name="file"
              title="Select File..."
              @file-update="update"
            >
              <template slot="icon-left">
                <v-icon left>add</v-icon>
              </template>
            </upload-btn>
            <v-btn
              block
              type="submit"
              color="success"
              @click.stop.prevent="submit()"
            >
              <v-icon class="mr-2">
                cloud_upload
              </v-icon>
              Submit
            </v-btn>
          </v-form>
          <v-progress-linear
            :indeterminate="true"
            :class="showBar"
          ></v-progress-linear>
          <v-alert
            :value="true"
            dismissible
            type="success"
            :class="showSuccess"
          >
            This is a success alert.
          </v-alert>
          <v-alert :value="true" dismissible type="error" :class="showError">
            This is a error alert.
          </v-alert>
        </v-card-text>
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script>
import UploadButton from 'vuetify-upload-button'

export default {
  components: {
    'upload-btn': UploadButton
  },
  data() {
    return {
      loading: false,
      file: '',
      success: false,
      error: false,
      errorMsg: ''
    }
  },
  computed: {
    showSuccess() {
      return this.success ? '' : 'd-none'
    },
    showError() {
      return this.error ? '' : 'd-none'
    },
    showBar() {
      return this.loading ? '' : 'd-none'
    }
  },
  methods: {
    update(file) {
      // handle file here. File will be an object.
      // If multiple prop is true, it will return an object array of files.
      this.file = file
    },
    submit() {
      this.loading = true
      /*
         Initialize the form data
      */
      const formData = new FormData()
      formData.append('file', this.file)

      this.$axios
        .$post('/aase/upload', formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        })
        .then(response => {
          this.loading = false
          this.success = true
          const url = window.URL.createObjectURL(new Blob([response.data]))
          const link = document.createElement('a')
          link.href = url
          link.setAttribute('download', 'fixed.apk')
          document.body.appendChild(link)
          link.click()
        })
        .catch(error => {
          this.errorMsg = error
          this.error = true
          this.loading = false
        })
    }
  }
}
</script>
