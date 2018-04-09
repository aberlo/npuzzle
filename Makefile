# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: ljoly <ljoly@student.42.fr>                +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2018/04/09 16:05:01 by ljoly             #+#    #+#              #
#    Updated: 2018/04/09 16:20:57 by ljoly            ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

NAME	=	npuzzle 
SRCS	=	$(shell find . -type f -name "*.go")

all: $(NAME)
	
$(NAME)	: $(SRCS)
	@go build -o $(NAME)
clean:
	@rm -f $(NAME)
re: clean all
run: all
	@./$(NAME) -f map/3_solvable.txt
