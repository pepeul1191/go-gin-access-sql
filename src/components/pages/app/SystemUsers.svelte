<script>
  export let id = null;

  // Lógica para cargar los usuarios del sistema, por ejemplo:
  let systemUsers = [];
  
  // Simular la carga de datos de usuarios (puedes usar fetch, por ejemplo)
  const fetchSystemUsers = async () => {
    // Esto es solo un ejemplo, deberías hacer una llamada a la API o base de datos
    systemUsers = [
      { username: "user1", email: "user1@example.com" },
      { username: "user2", email: "user2@example.com" }
    ];
  };

  // Cargar usuarios cuando el componente se monte
  fetchSystemUsers();
</script>

<h1>Usuarios del Sistema {id}</h1>

<!-- Mostrar los usuarios -->
<ul>
  {#each systemUsers as user}
    <li>{user.username} - {user.email}</li>
  {/each}
</ul>