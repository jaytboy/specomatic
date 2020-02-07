package main

//TODO:
//	1.	Extract zip files
//	2.	Get document.xml
//	3.	Parse xml
//	4.	Get <w:p...></w:p> and place in array
//	5.	Find all <w:pStyle w:val=""/> unique values
//	6.	Add value from above that represents the section title and number. This will be hard coded
//	7.	Create a json file like this:
//			[{<Section Number & Name>:[
//				{text:<test>,
//				level:<0>,
//				id:<id>,
//				parent:<reference id>},
//				{text:<test>,
//				level:<0>,
//				id:<id>,
//				parent:<reference id>},
//			]}]