<template>
  <v-container>

    <v-card
      class="mx-auto py-6 my-8 "
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
          >close</v-btn>
        </template>
      </v-snackbar>

      <form class="spacing-playground pa-6">
        <v-text-field
          v-model="$v.user.name.$model"
          label="姓名"
        ></v-text-field>
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
        <v-checkbox
          v-model="checkbox"
          :error-messages="checkboxErrors"
          label="同意条款?"
          required
          @change="$v.checkbox.$touch()"
          @blur="$v.checkbox.$touch()"
        ></v-checkbox>

        <v-btn
          color="success"
          @click="submit"
        >注册</v-btn>
        <v-btn
          color="error"
          @click="clear"
        >重置</v-btn>
      </form>
    </v-card>
  </v-container>
</template>
<script>
import { validationMixin } from 'vuelidate';
import { required, maxLength, minLength } from 'vuelidate/lib/validators';
import { mapActions } from 'vuex';

export default {
  data() {
    return {
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
      title: '注册',
      checkbox: false,
    };
  },

  mixins: [validationMixin],

  validations: {
    user: {
      name: '',
      phone: { required, maxLength: maxLength(11), minLength: minLength(11) },
      password: { required, maxLength: maxLength(15), minLength: minLength(6) },
    },
    checkbox: {
      checked(val) {
        return val;
      },
    },
  },

  computed: {
    checkboxErrors() {
      const errors = [];
      if (!this.$v.checkbox.$dirty) return errors;
      if (!this.$v.checkbox.checked) errors.push('必须同意才可以继续!');
      return errors;
    },
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
    ...mapActions('userModule', { userRegister: 'register' }),
    submit() {
      this.$v.$touch();
      if (this.$v.$anyError) {
        return;
      }
      this.userRegister(this.user).then((res) => {
        this.snackbar.text = res.data.data.msg;
        this.snackbar.show = true;
        this.$router.replace({ name: 'Home' });
      }).catch((err) => {
        this.snackbar.text = '注册失败';
        if (err.response) {
          this.snackbar.text = err.response.data.msg;
        }
        this.snackbar.show = true;
      });
    },
    clear() {
      this.$v.$reset();
      this.user.phone = '';
      this.user.password = '';
      this.user.checkbox = false;
    },
  },
};
</script>
<style lang="">
</style>
