function initialize() {

    var infoWindow;
        
    var MAPTYPEID = "custom_style";
    var bham = new google.maps.LatLng(33.6743890, -85.9455);


    var featureOpts = [
	{"featureType":"all","elementType":"all","stylers":[{"visibility":"on"}]},
	{"featureType":"all","elementType":"labels.text.fill","stylers":[{"color":"#040506"}]},
	{"featureType":"all","elementType":"labels.text.stroke","stylers":[{"color":"#FFFFFF"}]},
	{"featureType":"all","elementType":"labels.icon","stylers":[{"color":"#808080"},{"visibility":"off"}]},
	{"featureType":"landscape","elementType":"all","stylers":[{"color":"#3d3c3c"}]},
	{"featureType":"poi","elementType":"all","stylers":[{"color":"#f9ebec"}]},
	{"featureType":"poi.business","elementType":"all","stylers":[{"color":"#f8ebec"}]},
	{"featureType":"poi.park","elementType":"all","stylers":[{"color":"#f8ecec"}]},
	{"featureType":"poi.school","elementType":"all","stylers":[{"color":"#f9ebec"}]},
	{"featureType":"road.highway","elementType":"all","stylers":[{"color":"#983836"}]},
	{"featureType":"road.arterial","elementType":"all","stylers":[{"color":"#c05a58"}]},
	{"featureType":"road.local","elementType":"all","stylers":[{"color":"#e6bbbc"}]},
	{"featureType":"water","elementType":"all","stylers":[{"color":"#969696"}]}]; 
    var mapOptions = {
	center: bham, 
	zoom: 2,
	mapTypeControlOptions: {
	    mapTypeIds: [google.maps.MapTypeId.ROADMAP, MAPTYPEID]
	},
	mapTypeId: MAPTYPEID
    };
    

    var map = new google.maps.Map(document.getElementById('map-canvas'), mapOptions);
    var styledMapOptions = {
	name: 'Flog'
    };

    var customMapType = new google.maps.StyledMapType(featureOpts, styledMapOptions);
    

    map.mapTypes.set(MAPTYPEID, customMapType);

    var bounds = new google.maps.LatLngBounds();


    infoWindow = new google.maps.InfoWindow();
    console.log(events);
    for(var i = 0; i < events.length; i++){
	addMarker(events[i]);
    }

    function addMarker(event) {
	console.log(event);
	var loc = new google.maps.LatLng(event["Long"],event["Lat"]);
	console.log(loc);
        var marker = new google.maps.Marker({
            position: loc,
            map: map,
	    title: toString(i+1),
	    content: "<div class='jumbotron'><h3>"+event["Title"]+"</h3><p>"+event["Description"]+"</p></div>"
        });
	
	google.maps.event.addListener(marker, "click", function(){
	    showForm(marker);
	});
	$("#event-"+event["Id"]).click(function(){
	    showForm(marker);
	});

	bounds.extend(marker.position);
    }

    function showForm(marker) {
	if(infoWindow) infoWindow.close();
	infoWindow.setContent(marker.content);
	infoWindow.open(map, marker);
    }

    //google.maps.event.addListener(map, 'click', function(event) {
    //addMarker(event.latLng);
    //});

    map.fitBounds(bounds);

}

google.maps.event.addDomListener(window, 'load', initialize);
