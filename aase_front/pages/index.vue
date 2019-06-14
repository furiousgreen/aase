<template>
  <v-layout justify-center align-center>
    <v-flex xs12 sm8 md4>
      <v-card>
        <v-card-title class="headline">
          APK File
        </v-card-title>
        <v-card-text>
          <v-form>
            <upload-btn
              block
              fixed-width="100%"
              name="file"
              title="Select File..."
              @file-update="update"
            />
            <v-btn
              block
              type="submit"
              color="success"
              @click.stop.prevent="submit()"
            >
              Submit
            </v-btn>
          </v-form>
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
  methods: {
    update(file) {
      // handle file here. File will be an object.
      // If multiple prop is true, it will return an object array of files.
    }
  }
}
</script>
