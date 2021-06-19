<template>
  <v-card
    class="mx-auto py-6 my-10"
    max-width="500"
  >
    <v-card-title>
      <div class="text-h4 mx-auto">{{title}}</div>
    </v-card-title>
    <v-snackbar
      v-model="snackbar.show"
      :timeout="snackbar.timeout"
      top
    >
      {{ snackbar.text }}

      <template v-slot:action="{ attrs }">
        <v-btn
          color="blue"
          text
          v-bind="attrs"
          @click="snackbar.show = false"
        >
          close
        </v-btn>
      </template>
    </v-snackbar>
    <form class="spacing-playground pa-6">
      <v-text-field
        v-model="$v.user.phone.$model"
        :error-messages="phoneErrors"
        :counter="11"
        label="手机"
        required
        @input="$v.user.phone.$touch()"
        @blur="$v.user.phone.$touch()"
      ></v-text-field>
      <v-text-field
        v-model="$v.user.password.$model"
        :error-messages="passwordErrors"
        label="密码"
        type="password"
        required
        @input="$v.user.password.$touch()"
        @blur="$v.user.password.$touch()"
      ></v-text-field>

      <div class="mx-auto">
        <v-btn
          class="mx-auto"
          color="success"
          @click="submit"
        >登陆</v-btn>
        <v-btn
          class="mx-auto"
          color="error"
          @click="clear"
        >重置</v-btn>
      </div>
    </form>
  </v-card>
</template>
<script>
import { validationMixin } from 'vuelidate';
import { required, maxLength, minLength } from 'vuelidate/lib/validators';
import { mapActions } from 'vuex';

export default {
  data: () => ({
    user: {
      name: '',
      phone: '',
      password: '',
    },
    snackbar: {
      show: false,
      text: '',
      timeout: 2000,
    },
    title: '登陆',
  }),

  mixins: [validationMixin],

  validations: {
    user: {
      name: '',
      phone: { required, maxLength: maxLength(11), minLength: minLength(11) },
      password: { required, maxLength: maxLength(15), minLength: minLength(6) },
    },
  },

  computed: {
    phoneErrors() {
      const errors = [];
      if (!this.$v.user.phone.$dirty) return errors;
      if (!this.$v.user.phone.maxLength) errors.push('请输入有效手机号码.');
      if (!this.$v.user.phone.minLength) errors.push('请输入有效手机号码.');
      if (!this.$v.user.phone.required) errors.push('请输入手机号码.');
      return errors;
    },
    passwordErrors() {
      const errors = [];
      if (!this.$v.user.password.$dirty) return errors;
      if (!this.$v.user.password.maxLength) errors.push('密码至多15位。');
      if (!this.$v.user.password.minLength) errors.push('密码至少6位。');
      if (!this.$v.user.password.required) errors.push('请输入密码。');
      return errors;
    },
  },

  methods: {
    ...mapActions('userModule', { userLogin: 'login' }),
    submit() {
      this.$v.$touch();
      if (this.$v.$anyError) {
        return;
      }
      this.userLogin(this.user).then(() => {
        this.$router.replace({ name: 'Home' });
      }).catch((err) => {
        this.snackbar.text = '登陆失败!';
        if (err.response) {
          this.snackbar.text = err.data.data.msg;
        }
        this.snackbar.show = true;
      });
    },
    clear() {
      this.$v.$reset();
      this.user.phone = '';
      this.user.password = '';
    },
  },
};
</script>
<style lang="">
</style>
