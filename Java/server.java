import java.io.*;
import java.net.*;

class server
{
  public static void main(String argv[]) throws Exception
  {
    String clientSentence;
    String reversedSentence;
    ServerSocket welcomeSocket = new ServerSocket(6789);

    while(true)
    {
      Socket connectionSocket = welcomeSocket.accept();
      BufferedReader inFromClient = new BufferedReader(new InputStreamReader(connectionSocket.getInputStream()));
      DataOutputStream outToClient = new DataOutputStream(connectionSocket.getOutputStream());
      
      clientSentence = inFromClient.readLine();
      System.out.println("Received: " + clientSentence);
      reversedSentence = new StringBuilder(clientSentence).reverse().toString() + '\n';
      
      outToClient.writeBytes(reversedSentence);
    }
  }
}
