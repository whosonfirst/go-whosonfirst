window.addEventListener("load", function load(event){

    const null_island = [ 0.0, 0.0 ];
    const tile_url = "https://tile.openstreetmap.org/{z}/{x}/{y}.png";

    var record_layer;

    const record_layer_style = {
	radius: 8,
	fillColor: "#ff7800",
	color: "#000",
	weight: 1,
	opacity: 1,
	fillOpacity: 0.8
    };
     
    const parquet_el = document.querySelector("#parquet-uri");
    const example_parquet_el = document.querySelector("#example-parquet-uri");    
    const wof_el = document.querySelector("#wof-uri");
    const example_wof_el = document.querySelector("#example-wof-uri");    
    const submit_el = document.querySelector("#find-record");
    const geojson_el = document.querySelector("#geojson");
    const map_el = document.querySelector("#map");        
    
    sfomuseum.golang.wasm.fetch("wasm/parquet_find_record.wasm").then((rsp) => {

	submit_el.removeAttribute("disabled");

	example_parquet_el.onclick = function(){
	    parquet_el.value = example_parquet_el.innerText;
	    return false;
	};

	example_wof_el.onclick = function(){
	    wof_el.value = example_wof_el.innerText;
	    return false;
	};

	const map = L.map(map_el.getAttribute("id"));

	map.setView(null_island, 1);
	
        var tile_layer = L.tileLayer(tile_url, {
            maxZoom: 19,
	    attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
        });
	
        tile_layer.addTo(map);
	
	submit_el.onclick = function(){

	    const parquet_uri = parquet_el.value;
	    const wof_uri = wof_el.value;
	    
	    const u = new URL(location);
	    u.pathname = "/parquet/" + parquet_uri;

	    geojson_el.innerHTML = "";
	    
	    parquet_find_record(u.toString(), wof_uri).then((rsp) => {
		return JSON.parse(rsp);
	    }).then((record) => {
		console.log(record);

		const enc = JSON.stringify(record, null, 2);

		const pre = document.createElement("pre");
		pre.appendChild(document.createTextNode(enc));
		geojson.appendChild(pre);

		if (record_layer){
		    map.removeLayer(record_layer);
		}

		const bounds = whosonfirst.geojson.deriveBboxAsBounds(record);
		map.fitBounds(bounds);

		const record_args = {
		    pointToLayer: function (feature, latlng) {
			return L.circleMarker(latlng, record_layer_style);
		    }
		};
		
		record_layer = L.geoJSON(record, record_args);
		record_layer.addTo(map);
		
	    }).catch((err) => {
		console.error("Failed to find record", err);
	    });

	    return false;
	};
	
    }).catch((err) => {
	console.error("Failed to load update exif binary", err);
        return;
    });
    
});
