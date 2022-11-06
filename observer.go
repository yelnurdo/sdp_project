class Messenger extends PVPatterns {
    public function send($message, $user) {
 
      $data = array(
        ‘message’ => $message,
        ‘first_name’ => $user -> first_name,
        ‘last_name’ => $user -> last_name
      );
 
     $result = Slack::sendMessage($data);
     //Call observers that are instances
     self::_notify(get_class() . ‘::’ . __FUNCTION__, $this,$message, $user, $result);
     //Calls observers that are static
     self::_notify(get_called_class() . ‘::’ . __FUNCTION__, $this,     $message, $user, $result);
    }
}

class Messenger extends PVPatterns {
	public function send($message, $user) {
   
	   $data = array(
		 ‘message’ => $message,
		 ‘first_name’ => $user -> first_name,
		 ‘last_name’ => $user -> last_name
	   );
   
	   Slack::sendMessage($data);
	}
  }