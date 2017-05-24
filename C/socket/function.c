#include <sys/socket.h>
  int socket (int family, int type, int protocol);

#include <sys/socket.h>
  int connect (int sockfd, const struct sockaddr *servaddr, socklen_t addrlen);
  
#include <sys/socket.h>
  int bind(int sockfd, const struct sockaddr *servaddr, socklen_t addrlen);
  
struct sockaddr_in serv; /* IPv4 socket address structure */
  bind(sockfd, (struct sockaddr*) &serv, sizeof(serv))
  
#include <sys/socket.h>
  int listen(int sockfd, int backlog);
  
#include <sys/socket.h>
  int accept(int sockfd, struct sockaddr *cliaddr, socklen_t *addrlen);
