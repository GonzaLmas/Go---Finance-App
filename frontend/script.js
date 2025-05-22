document.querySelectorAll('.grid-button').forEach(button => {
    button.addEventListener('click', function() {
        const tablaContainer = document.getElementById('tablaResultados');
        const loadingDiv = document.createElement('div');
        loadingDiv.className = 'loading';
        loadingDiv.textContent = 'Cargando datos...';
        loadingDiv.id = 'loading-' + this.id.replace('btn', '').toLowerCase();
        tablaContainer.appendChild(loadingDiv);
        
        const endpoint = '/' + this.id.replace('btn', '').toLowerCase();
        
        fetch(endpoint)
            .then(response => {
                if (!response.ok) throw new Error(`Error ${response.status}`);
                return response.json();
            })
            .then(data => {
                if (!data) throw new Error("No se recibieron datos");
                // Eliminar el mensaje de carga correspondiente
                const loadingToRemove = document.getElementById('loading-' + this.id.replace('btn', '').toLowerCase());
                if (loadingToRemove) loadingToRemove.remove();
                mostrarTablaInstrumentos(data, this.textContent.trim());
            })
            .catch(error => {
                // Reemplazar el mensaje de carga con el error
                const loadingToReplace = document.getElementById('loading-' + this.id.replace('btn', '').toLowerCase());
                if (loadingToReplace) {
                    loadingToReplace.className = 'error';
                    loadingToReplace.textContent = `Error: ${error.message}`;
                }
                console.error('Error:', error);
            });
    });
});

function mostrarTablaInstrumentos(data, titulo) {
    const tablaContainer = document.getElementById('tablaResultados');
    
    // Crear contenedor de tabla
    const tablaWrapper = document.createElement('div');
    tablaWrapper.className = 'tabla-wrapper';
    
    // Añadir título
    const tituloElement = document.createElement('h3');
    tituloElement.className = 'tabla-titulo';
    tituloElement.textContent = titulo;
    tablaWrapper.appendChild(tituloElement);
    
    // Crear tabla
    const table = document.createElement('table');
    table.className = 'data-table';
    
    // Crear encabezados
    const thead = document.createElement('thead');
    const headerRow = document.createElement('tr');
    
    ['Instrumento', 'Precio'].forEach(texto => {
        const th = document.createElement('th');
        th.textContent = texto;
        headerRow.appendChild(th);
    });
    
    thead.appendChild(headerRow);
    table.appendChild(thead);
    
    // Llenar datos
    const tbody = document.createElement('tbody');
    const itemsMostrar = Array.isArray(data) ? data.slice(0, 10) : [];
    
    itemsMostrar.forEach(item => {
        const row = document.createElement('tr');
        
        // Celda Instrumento
        const tdInstrumento = document.createElement('td');
        tdInstrumento.textContent = item.symbol || item.ticker || item.instrumento || 'N/A';
        row.appendChild(tdInstrumento);
        
        // Celda Precio
        const tdPrecio = document.createElement('td');
        tdPrecio.className = 'price-cell';
        let precio = item.c || item.price || item.precio || item.value;
        tdPrecio.textContent = typeof precio === 'number' ? `$${precio.toFixed(2)}` : precio || 'N/A';
        row.appendChild(tdPrecio);
        
        tbody.appendChild(row);
    });
    
    // Si hay más items, crear fila clickeable
    if (Array.isArray(data) && data.length > 10) {
        const row = document.createElement('tr');
        row.className = 'ver-mas-row';
        row.style.cursor = 'pointer';
        
        const td = document.createElement('td');
        td.colSpan = 2;
        td.style.textAlign = 'center';
        td.style.padding = '10px';
        
        const verMasBtn = document.createElement('span');
        verMasBtn.className = 'ver-mas-btn';
        verMasBtn.textContent = `Ver ${data.length - 10} más...`;
        verMasBtn.style.color = '#6f42c1';
        verMasBtn.style.textDecoration = 'underline';
        verMasBtn.style.fontWeight = '600';
        
        td.appendChild(verMasBtn);
        row.appendChild(td);
        
        // Evento para mostrar todos los items
        row.addEventListener('click', () => {
            // Eliminar la fila "Ver más"
            row.remove();
            
            // Mostrar todos los items restantes
            data.slice(10).forEach(item => {
                const newRow = document.createElement('tr');
                
                const tdInstrumento = document.createElement('td');
                tdInstrumento.textContent = item.symbol || item.ticker || item.instrumento || 'N/A';
                newRow.appendChild(tdInstrumento);
                
                const tdPrecio = document.createElement('td');
                tdPrecio.className = 'price-cell';
                let precio = item.c || item.price || item.precio || item.value;
                tdPrecio.textContent = typeof precio === 'number' ? `$${precio.toFixed(2)}` : precio || 'N/A';
                newRow.appendChild(tdPrecio);
                
                tbody.appendChild(newRow);
            });
        });
        
        tbody.appendChild(row);
    }
    
    table.appendChild(tbody);
    tablaWrapper.appendChild(table);
    tablaContainer.insertBefore(tablaWrapper, tablaContainer.firstChild);
    
    // Limitar a 9 tablas (3 filas)
    const tablas = tablaContainer.querySelectorAll('.tabla-wrapper');
    if (tablas.length > 9) {
        tablaContainer.removeChild(tablas[tablas.length - 1]);
    }
}