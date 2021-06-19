<template lang="">
<div>
    <div>
        <!-- footer -->
        <v-footer absolute class="font-weight-medium">
            <v-col class="text-center" cols="12">
                {{ new Date().getFullYear() }} — <strong>{{title}}</strong>
            </v-col>
        </v-footer>
    </div>
</div>
</template>
<script>
export default {
  data() {
    return {
      title: 'min',
    };
  },
};
</script>
<style lang="">
</style>
