document.querySelectorAll('.grid-button').forEach(button => {
    button.addEventListener('click', function() {
        const tablaContainer = document.getElementById('tablaResultados');
        const buttonId = this.id.replace('btn', '').toLowerCase();
        const loadingId = `loading-${buttonId}`;
        const buttonText = this.textContent.trim();
        const tablaId = `tabla-${buttonId}`;
        
        const tablasExistentes = document.querySelectorAll('.tabla-wrapper');
        let tablaExistente = null;
        
        tablasExistentes.forEach(tabla => {
            const titulo = tabla.querySelector('h3.tabla-titulo');
            if (titulo && titulo.textContent.trim() === buttonText) {
                tablaExistente = tabla;
            }
        });
        
        if (tablaExistente) {
            tablaContainer.removeChild(tablaExistente);
        }
        
        const loadingDiv = document.createElement('div');
        loadingDiv.className = 'loading';
        loadingDiv.textContent = 'Cargando datos...';
        loadingDiv.id = loadingId;
        tablaContainer.appendChild(loadingDiv);
        
        const endpoint = `/${buttonId}`;
        
        fetch(endpoint)
            .then(response => {
                if (!response.ok) throw new Error(`Error ${response.status}`);
                return response.json();
            })
            .then(data => {
                if (!data) throw new Error("No se recibieron datos");
                
                // Eliminar loading
                const loadingToRemove = document.getElementById(loadingId);
                if (loadingToRemove) loadingToRemove.remove();
                
                // Crear nueva tabla (siempre)
                if (buttonText.includes('DÓLAR') && !buttonText.includes('HISTÓRICO')) {
                    crearTablaDolar(data, buttonText, tablaId);
                } else if (buttonText.includes('HISTÓRICO')) {
                    crearTablaDolarHistorico(data, buttonText, tablaId);
                } else if (endpoint.includes('/coinbase')) {
                    crearTablaCoinbase(data, buttonText, tablaId);
                } else {
                    crearTablaInstrumentos(data, buttonText, tablaId);
                }
                
                const tablas = document.querySelectorAll('.tabla-wrapper');
                if (tablas.length > 9) {
                    tablaContainer.removeChild(tablas[tablas.length - 1]);
                }
            })
            .catch(error => {
                const loadingToReplace = document.getElementById(loadingId);
                if (loadingToReplace) {
                    loadingToReplace.className = 'error';
                    loadingToReplace.textContent = `Error: ${error.message}`;
                }
                console.error('Error:', error);
            });
    });
});

function crearTablaInstrumentos(data, titulo, tablaId) {
    const tablaContainer = document.getElementById('tablaResultados');
    const tablaWrapper = document.createElement('div');
    tablaWrapper.className = 'tabla-wrapper';
    tablaWrapper.id = tablaId;
    
    const tituloElement = document.createElement('h3');
    tituloElement.className = 'tabla-titulo';
    tituloElement.textContent = titulo;
    tablaWrapper.appendChild(tituloElement);
    
    const tableContainer = document.createElement('div');
    tableContainer.className = 'data-table-container';
    
    const table = document.createElement('table');
    table.className = 'data-table';
    
    const thead = document.createElement('thead');
    const headerRow = document.createElement('tr');
    ['Instrumento', 'Precio'].forEach(texto => {
        const th = document.createElement('th');
        th.textContent = texto;
        headerRow.appendChild(th);
    });
    thead.appendChild(headerRow);
    table.appendChild(thead);
    
    const tbody = document.createElement('tbody');
    const datos = Array.isArray(data) ? data : [data];
    
    datos.slice(0, 10).forEach(item => {
        const row = document.createElement('tr');
        
        const tdInstrumento = document.createElement('td');
        tdInstrumento.textContent = item.symbol || item.ticker || item.instrumento || 'N/A';
        row.appendChild(tdInstrumento);
        
        const tdPrecio = document.createElement('td');
        tdPrecio.className = 'price-cell';
        const precio = item.c || item.price || item.precio || item.value;
        tdPrecio.textContent = precio !== undefined ? `$${parseFloat(precio).toFixed(2)}` : 'N/A';
        row.appendChild(tdPrecio);
        
        tbody.appendChild(row);
    });
    
    if (datos.length > 10) {
        const row = document.createElement('tr');
        row.className = 'ver-mas-row';
        const td = document.createElement('td');
        td.colSpan = 2;
        td.style.textAlign = 'center';
        
        const verMasBtn = document.createElement('span');
        verMasBtn.className = 'ver-mas-btn';
        verMasBtn.textContent = `Ver ${datos.length - 10} más...`;
        verMasBtn.onclick = () => {
            row.remove();
            datos.slice(10).forEach(item => {
                const newRow = document.createElement('tr');
                
                const tdInstrumento = document.createElement('td');
                tdInstrumento.textContent = item.symbol || item.ticker || item.instrumento || 'N/A';
                newRow.appendChild(tdInstrumento);
                
                const tdPrecio = document.createElement('td');
                tdPrecio.className = 'price-cell';
                const precio = item.c || item.price || item.precio || item.value;
                tdPrecio.textContent = precio !== undefined ? `$${parseFloat(precio).toFixed(2)}` : 'N/A';
                newRow.appendChild(tdPrecio);
                
                tbody.appendChild(newRow);
            });
        };
        
        td.appendChild(verMasBtn);
        row.appendChild(td);
        tbody.appendChild(row);
    }
    
    table.appendChild(tbody);
    tableContainer.appendChild(table);
    tablaWrapper.appendChild(tableContainer);
    tablaContainer.insertBefore(tablaWrapper, tablaContainer.firstChild);
}

function crearTablaDolar(data, titulo, tablaId) {
    const tablaContainer = document.getElementById('tablaResultados');
    const tablaWrapper = document.createElement('div');
    tablaWrapper.className = 'tabla-wrapper';
    tablaWrapper.id = tablaId;
    
    const tituloElement = document.createElement('h3');
    tituloElement.className = 'tabla-titulo';
    tituloElement.textContent = titulo;
    tablaWrapper.appendChild(tituloElement);
    
    const tableContainer = document.createElement('div');
    tableContainer.className = 'data-table-container';
    
    const table = document.createElement('table');
    table.className = 'data-table';
    
    const thead = document.createElement('thead');
    const headerRow = document.createElement('tr');
    ['Tipo', 'Compra', 'Venta'].forEach(texto => {
        const th = document.createElement('th');
        th.textContent = texto;
        headerRow.appendChild(th);
    });
    thead.appendChild(headerRow);
    table.appendChild(thead);
    
    const tbody = document.createElement('tbody');
    const datos = Array.isArray(data) ? data : [data];
    
    datos.forEach(item => {
        const row = document.createElement('tr');
        
        const tdTipo = document.createElement('td');
        tdTipo.textContent = item.nombre || item.casa || 'N/A';
        row.appendChild(tdTipo);
        
        const tdCompra = document.createElement('td');
        tdCompra.className = 'price-cell';
        tdCompra.textContent = item.compra !== undefined ? `$${item.compra.toFixed(2)}` : 'N/A';
        row.appendChild(tdCompra);
        
        const tdVenta = document.createElement('td');
        tdVenta.className = 'price-cell';
        tdVenta.textContent = item.venta !== undefined ? `$${item.venta.toFixed(2)}` : 'N/A';
        row.appendChild(tdVenta);
        
        tbody.appendChild(row);
    });
    
    table.appendChild(tbody);
    tableContainer.appendChild(table);
    tablaWrapper.appendChild(tableContainer);
    tablaContainer.insertBefore(tablaWrapper, tablaContainer.firstChild);
}

function crearTablaCoinbase(data, titulo, tablaId) {
    const tablaContainer = document.getElementById('tablaResultados');
    const tablaWrapper = document.createElement('div');
    tablaWrapper.className = 'tabla-wrapper';
    tablaWrapper.id = tablaId;
    
    // Resto del código de la función original...
    const tituloElement = document.createElement('h3');
    tituloElement.className = 'tabla-titulo';
    tituloElement.textContent = titulo;
    tablaWrapper.appendChild(tituloElement);
    
    const table = document.createElement('table');
    table.className = 'data-table';
    
    const thead = document.createElement('thead');
    const headerRow = document.createElement('tr');
    ['Cripto', 'Precio (USD)'].forEach(texto => {
        const th = document.createElement('th');
        th.textContent = texto;
        headerRow.appendChild(th);
    });
    thead.appendChild(headerRow);
    table.appendChild(thead);
    
    const tbody = document.createElement('tbody');
    const datos = Array.isArray(data) ? data : [data];
    
    datos.forEach(item => {
        const row = document.createElement('tr');
        
        const tdCripto = document.createElement('td');
        tdCripto.textContent = item.base || 'N/A';
        row.appendChild(tdCripto);
        
        const tdPrecio = document.createElement('td');
        tdPrecio.className = 'price-cell';
        const precio = item.amount ? `$${parseFloat(item.amount).toFixed(2)}` : 'N/A';
        tdPrecio.textContent = precio;
        row.appendChild(tdPrecio);
        
        tbody.appendChild(row);
    });
    
    table.appendChild(tbody);
    tablaWrapper.appendChild(table);
    tablaContainer.insertBefore(tablaWrapper, tablaContainer.firstChild);
}