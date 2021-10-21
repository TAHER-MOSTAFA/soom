var log = msg => {
    document.getElementById('logs').innerHTML += msg + '<br>'
  }
let pc = new RTCPeerConnection({
    iceServers: [
      {
        urls: 'stun:stun.l.google.com:19302'
      }
    ]
  })

pc.oniceconnectionstatechange = e => log(pc.iceConnectionState)
pc.onicecandidate = event => {
    if (event.candidate === null) {
        var name = document.getElementById("name").value
        var offer =  btoa(JSON.stringify(pc.localDescription))
       
        answer = fetch(location.pathname, {
            method: "POST", 
            body: JSON.stringify({name : name, offer:offer})
          }).then(res => {
            res.json().then(data => {return data});
          });
        
        try{
        pc.setRemoteDescription(new RTCSessionDescription(JSON.parse(atob(data["answer"]))))
        } catch (e) {
            alert(e)
        }
    }    
}

function startmeeting() {
    var name = document.getElementById("name").value
    if (name === ""){
        return alert("u can't join with blank name")
    }   
    navigator.mediaDevices.getUserMedia({ video: true, audio: true })
      .then(stream => {
        stream.getTracks().forEach(track => pc.addTrack(track, stream));
        document.getElementById('video').srcObject = stream
       
        pc.createOffer()
          .then(d => pc.setLocalDescription(d))
          .catch(log)

      }).catch(log)
}