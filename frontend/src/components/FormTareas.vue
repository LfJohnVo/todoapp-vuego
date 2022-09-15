<template>
  <div id="form-tareas">
    <form class="d-flex flex-row align-items-center flex-wrap">
      <div class="form-check my-1 me-sm-4">
        <input type="text" id="tarea" class="form-control" placeholder="Ingrese tarea a cargar ..." v-model="tarea" v-validate="'required'"/>
      </div>
      <button type="submit" class="btn btn-primary my-1">Enviar</button>
    </form>
  </div>
</template>
    
    <script>
export default {
  name: "form-tareas",
};
</script>
    
    <style scoped></style>