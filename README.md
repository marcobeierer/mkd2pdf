# MKD2PDF

Elementary command line frontend/assistant for the universal document converter [http://johnmacfarlane.net/pandoc/](Pandoc) to convert Markdown files to the PDF.

The intent of the program is to provide a easy way to use Pandoc for users, which are not confident with the command line.

## Prerequisites
- Pandoc

## Usage
The program could be executed directly on the command line:

	./mkd2pdf

However for the target audience of this program it would be easier to include the program in the context menu of the used file manager. In Thunar, the file manager of XFCE4, you could use the following command in a custom context menu entry:
	
	xfce4-terminal --working-directory %f --hold -x /path/to/mkd2pdf
