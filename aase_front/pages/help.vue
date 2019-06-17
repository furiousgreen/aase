<template>
  <v-layout align-start justify-start>
    <v-flex>
      <v-card class="ma-2">
        <v-card-title class="primary white--text primary-title headline">
          Help
        </v-card-title>
        <v-card-text class="pa-5">
          <ul>
            <li class="title">
              Quickstart
            </li>
            <p class="body-1 mt-2">
              Choose the APK you'd like to modify by clicking the "Select
              File..." button. After selecting the file click on "Submit".
            </p>
          </ul>
        </v-card-text>
      </v-card>
      <v-card class="ma-2">
        <v-card-title class="primary white--text primary-title headline">
          Frequently Asked Questions
        </v-card-title>
        <v-card-text class="pa-5">
          <ul>
            <li class="title">
              What is this Android Security Exception?
            </li>
            <p class="body-1 mt-2">
              Since Android 7.0, Android only allows certificates specified
              inside app, in the netwok_security_config.xml. To allow Charles to
              look at the contents of the network connections, we need to add
              Charles to the list of trusted certificates.
            </p>
            <p>
              You can find more about it
              <a
                href="https://developer.android.com/training/articles/security-config"
              >
                here
              </a>
            </p>
            <li class="title">
              How are you adding the exception?
            </li>
            <p class="body-1 mt-2">
              This tool is a web interface for the tool built by
              <a href="https://github.com/levyitay/">Levy Itay</a>. If you find
              the source code
              <a href="https://github.com/levyitay/AddSecurityExceptionAndroid">
                here</a
              >.
            </p>
          </ul>
        </v-card-text>
      </v-card>
    </v-flex>
  </v-layout>
</template>
