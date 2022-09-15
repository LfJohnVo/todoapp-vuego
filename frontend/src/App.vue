<template>
  <div class="container">
    <form-tareas />
    <div id="app" class="container">
      <div class="row">
        <div class="col-md-12">
          <h1>Tareas</h1>
        </div>
      </div>
      <div class="row">
        <div class="col-md-12">
          <tabla-tareas />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import TablaTareas from "@/components/TablaTareas.vue";
import FormTareas from "@/components/FormTareas.vue";

export default {
  name: "app",
  components: {
    TablaTareas,
    FormTareas,
  },
};
</script>
  
<style>
button {
  background: #009435;
  border: 1px solid #009435;
}
</style>