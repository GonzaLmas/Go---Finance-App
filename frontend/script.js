document.querySelectorAll('.grid-button').forEach(button => {
    button.addEventListener('click', function() {
        const tablaContainer = document.getElementById('tablaResultados');
        tablaContainer.innerHTML = '<div class="loading">Cargando datos...</div>';
        
        const endpoint = '/' + this.id.replace('btn', '').toLowerCase();
        
        fetch(endpoint)
            .then(response => {
                if (!response.ok) throw new Error(`Error ${response.status}`);
                return response.json();
            })
            .then(data => {
                if (!data) throw new Error("No se recibieron datos");
                mostrarTablaInstrumentos(data);
            })
            .catch(error => {
                tablaContainer.innerHTML = `<div class="error">Error: ${error.message}</div>`;
                console.error('Error:', error);
            });
    });
});

function mostrarTablaInstrumentos(data) {
    const tablaContainer = document.getElementById('tablaResultados');
    tablaContainer.innerHTML = '';
    
    if (!data || (Array.isArray(data) && data.length === 0)) {
        tablaContainer.innerHTML = '<div class="error">No hay datos disponibles</div>';
        return;
    }

    // Normalizar datos (array o objeto)
    const datos = Array.isArray(data) ? data : [data];
    
    // Crear tabla
    const table = document.createElement('table');
    table.className = 'data-table';
    
    // Encabezados fijos
    const thead = document.createElement('thead');
    const headerRow = document.createElement('tr');
    
    ['Instrumento', 'Precio'].forEach(texto => {
        const th = document.createElement('th');
        th.textContent = texto;
        headerRow.appendChild(th);
    });
    
    thead.appendChild(headerRow);
    table.appendChild(thead);
    
    // Cuerpo de la tabla
    const tbody = document.createElement('tbody');
    
    datos.forEach(item => {
        const row = document.createElement('tr');
        
        // Celda Instrumento (usa symbol, ticker o el primer campo disponible)
        const instrumento = item.symbol || item.ticker || Object.values(item)[0];
        const tdInstrumento = document.createElement('td');
        tdInstrumento.textContent = instrumento || 'N/A';
        row.appendChild(tdInstrumento);
        
        // Celda Precio (usa c, price, precio o el segundo campo disponible)
        let precio = item.c || item.price || item.precio || Object.values(item)[1];
        const tdPrecio = document.createElement('td');
        tdPrecio.className = 'price-cell';
        
        if (typeof precio === 'number') {
            // Formatear como $11.43
            precio = '$' + precio.toFixed(2).replace(/\B(?=(\d{3})+(?!\d))/g, ",");
        }
        tdPrecio.textContent = precio !== undefined ? precio : 'N/A';
        row.appendChild(tdPrecio);
        
        tbody.appendChild(row);
    });
    
    table.appendChild(tbody);
    tablaContainer.appendChild(table);
}