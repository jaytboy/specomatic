package main

// Body accepts a document paragraph structure
type Body struct {
	Text string `xml:",chardata"`
	P    []struct {
		Text         string `xml:",chardata"`
		RsidR        string `xml:"rsidR,attr"`
		RsidRPr      string `xml:"rsidRPr,attr"`
		RsidRDefault string `xml:"rsidRDefault,attr"`
		RsidP        string `xml:"rsidP,attr"`
		PPr          struct {
			Text   string `xml:",chardata"`
			PStyle struct {
				Text string `xml:",chardata"`
				Val  string `xml:"val,attr"`
			} `xml:"pStyle"`
			NumPr struct {
				Text string `xml:",chardata"`
				Ilvl struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
				} `xml:"ilvl"`
				NumID struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
				} `xml:"numId"`
			} `xml:"numPr"`
			Ind struct {
				Text    string `xml:",chardata"`
				Left    string `xml:"left,attr"`
				Hanging string `xml:"hanging,attr"`
			} `xml:"ind"`
			RPr struct {
				Text   string `xml:",chardata"`
				RStyle struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
				} `xml:"rStyle"`
				RFonts struct {
					Text string `xml:",chardata"`
					Cs   string `xml:"cs,attr"`
				} `xml:"rFonts"`
				Highlight struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
				} `xml:"highlight"`
				Sz struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
				} `xml:"sz"`
			} `xml:"rPr"`
			Tabs struct {
				Text string `xml:",chardata"`
				Tab  []struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
					Pos  string `xml:"pos,attr"`
				} `xml:"tab"`
			} `xml:"tabs"`
			KeepNext struct {
				Text string `xml:",chardata"`
				Val  string `xml:"val,attr"`
			} `xml:"keepNext"`
		} `xml:"pPr"`
		R []struct {
			Text    string `xml:",chardata"`
			RsidRPr string `xml:"rsidRPr,attr"`
			RsidR   string `xml:"rsidR,attr"`
			T       struct {
				Text  string `xml:",chardata"`
				Space string `xml:"space,attr"`
			} `xml:"t"`
			RPr struct {
				Text string `xml:",chardata"`
				B    string `xml:"b"`
				U    struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
				} `xml:"u"`
				VertAlign struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
				} `xml:"vertAlign"`
				RStyle struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
				} `xml:"rStyle"`
				Highlight struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
				} `xml:"highlight"`
			} `xml:"rPr"`
			LastRenderedPageBreak string `xml:"lastRenderedPageBreak"`
		} `xml:"r"`
	} `xml:"p"`
	Tbl []struct {
		Text  string `xml:",chardata"`
		TblPr struct {
			Text string `xml:",chardata"`
			TblW struct {
				Text string `xml:",chardata"`
				W    string `xml:"w,attr"`
				Type string `xml:"type,attr"`
			} `xml:"tblW"`
			TblBorders struct {
				Text string `xml:",chardata"`
				Top  struct {
					Text  string `xml:",chardata"`
					Val   string `xml:"val,attr"`
					Sz    string `xml:"sz,attr"`
					Space string `xml:"space,attr"`
					Color string `xml:"color,attr"`
				} `xml:"top"`
				Left struct {
					Text  string `xml:",chardata"`
					Val   string `xml:"val,attr"`
					Sz    string `xml:"sz,attr"`
					Space string `xml:"space,attr"`
					Color string `xml:"color,attr"`
				} `xml:"left"`
				Bottom struct {
					Text  string `xml:",chardata"`
					Val   string `xml:"val,attr"`
					Sz    string `xml:"sz,attr"`
					Space string `xml:"space,attr"`
					Color string `xml:"color,attr"`
				} `xml:"bottom"`
				Right struct {
					Text  string `xml:",chardata"`
					Val   string `xml:"val,attr"`
					Sz    string `xml:"sz,attr"`
					Space string `xml:"space,attr"`
					Color string `xml:"color,attr"`
				} `xml:"right"`
				InsideH struct {
					Text  string `xml:",chardata"`
					Val   string `xml:"val,attr"`
					Sz    string `xml:"sz,attr"`
					Space string `xml:"space,attr"`
					Color string `xml:"color,attr"`
				} `xml:"insideH"`
				InsideV struct {
					Text  string `xml:",chardata"`
					Val   string `xml:"val,attr"`
					Sz    string `xml:"sz,attr"`
					Space string `xml:"space,attr"`
					Color string `xml:"color,attr"`
				} `xml:"insideV"`
			} `xml:"tblBorders"`
			TblLook struct {
				Text        string `xml:",chardata"`
				Val         string `xml:"val,attr"`
				FirstRow    string `xml:"firstRow,attr"`
				LastRow     string `xml:"lastRow,attr"`
				FirstColumn string `xml:"firstColumn,attr"`
				LastColumn  string `xml:"lastColumn,attr"`
				NoHBand     string `xml:"noHBand,attr"`
				NoVBand     string `xml:"noVBand,attr"`
			} `xml:"tblLook"`
			TblInd struct {
				Text string `xml:",chardata"`
				W    string `xml:"w,attr"`
				Type string `xml:"type,attr"`
			} `xml:"tblInd"`
			TblLayout struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"tblLayout"`
		} `xml:"tblPr"`
		TblGrid struct {
			Text    string `xml:",chardata"`
			GridCol []struct {
				Text string `xml:",chardata"`
				W    string `xml:"w,attr"`
			} `xml:"gridCol"`
		} `xml:"tblGrid"`
		Tr []struct {
			Text   string `xml:",chardata"`
			RsidR  string `xml:"rsidR,attr"`
			RsidTr string `xml:"rsidTr,attr"`
			TrPr   struct {
				Text      string `xml:",chardata"`
				CantSplit string `xml:"cantSplit"`
				TrHeight  struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
				} `xml:"trHeight"`
				TblHeader string `xml:"tblHeader"`
			} `xml:"trPr"`
			Tc []struct {
				Text string `xml:",chardata"`
				TcPr struct {
					Text string `xml:",chardata"`
					TcW  struct {
						Text string `xml:",chardata"`
						W    string `xml:"w,attr"`
						Type string `xml:"type,attr"`
					} `xml:"tcW"`
					TcBorders struct {
						Text string `xml:",chardata"`
						Top  struct {
							Text  string `xml:",chardata"`
							Val   string `xml:"val,attr"`
							Sz    string `xml:"sz,attr"`
							Space string `xml:"space,attr"`
							Color string `xml:"color,attr"`
						} `xml:"top"`
						Left struct {
							Text  string `xml:",chardata"`
							Val   string `xml:"val,attr"`
							Sz    string `xml:"sz,attr"`
							Space string `xml:"space,attr"`
							Color string `xml:"color,attr"`
						} `xml:"left"`
						Bottom struct {
							Text  string `xml:",chardata"`
							Val   string `xml:"val,attr"`
							Sz    string `xml:"sz,attr"`
							Space string `xml:"space,attr"`
							Color string `xml:"color,attr"`
						} `xml:"bottom"`
						Right struct {
							Text  string `xml:",chardata"`
							Val   string `xml:"val,attr"`
							Sz    string `xml:"sz,attr"`
							Space string `xml:"space,attr"`
							Color string `xml:"color,attr"`
						} `xml:"right"`
					} `xml:"tcBorders"`
					HideMark string `xml:"hideMark"`
					VAlign   struct {
						Text string `xml:",chardata"`
						Val  string `xml:"val,attr"`
					} `xml:"vAlign"`
				} `xml:"tcPr"`
				P struct {
					Text         string `xml:",chardata"`
					RsidR        string `xml:"rsidR,attr"`
					RsidRPr      string `xml:"rsidRPr,attr"`
					RsidRDefault string `xml:"rsidRDefault,attr"`
					RsidP        string `xml:"rsidP,attr"`
					PPr          struct {
						Text   string `xml:",chardata"`
						PStyle struct {
							Text string `xml:",chardata"`
							Val  string `xml:"val,attr"`
						} `xml:"pStyle"`
						Spacing struct {
							Text   string `xml:",chardata"`
							Before string `xml:"before,attr"`
							After  string `xml:"after,attr"`
						} `xml:"spacing"`
						Ind struct {
							Text string `xml:",chardata"`
							Left string `xml:"left,attr"`
						} `xml:"ind"`
						RPr struct {
							Text string `xml:",chardata"`
							B    string `xml:"b"`
							Sz   struct {
								Text string `xml:",chardata"`
								Val  string `xml:"val,attr"`
							} `xml:"sz"`
							SzCs struct {
								Text string `xml:",chardata"`
								Val  string `xml:"val,attr"`
							} `xml:"szCs"`
						} `xml:"rPr"`
					} `xml:"pPr"`
					R []struct {
						Text    string `xml:",chardata"`
						RsidRPr string `xml:"rsidRPr,attr"`
						RsidR   string `xml:"rsidR,attr"`
						RPr     struct {
							Text string `xml:",chardata"`
							B    string `xml:"b"`
						} `xml:"rPr"`
						T struct {
							Text  string `xml:",chardata"`
							Space string `xml:"space,attr"`
						} `xml:"t"`
					} `xml:"r"`
				} `xml:"p"`
			} `xml:"tc"`
		} `xml:"tr"`
	} `xml:"tbl"`
	SectPr struct {
		Text            string `xml:",chardata"`
		RsidR           string `xml:"rsidR,attr"`
		RsidRPr         string `xml:"rsidRPr,attr"`
		RsidSect        string `xml:"rsidSect,attr"`
		HeaderReference []struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			ID   string `xml:"id,attr"`
		} `xml:"headerReference"`
		FooterReference []struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			ID   string `xml:"id,attr"`
		} `xml:"footerReference"`
		FootnotePr struct {
			Text       string `xml:",chardata"`
			NumRestart struct {
				Text string `xml:",chardata"`
				Val  string `xml:"val,attr"`
			} `xml:"numRestart"`
		} `xml:"footnotePr"`
		EndnotePr struct {
			Text   string `xml:",chardata"`
			NumFmt struct {
				Text string `xml:",chardata"`
				Val  string `xml:"val,attr"`
			} `xml:"numFmt"`
		} `xml:"endnotePr"`
		PgSz struct {
			Text string `xml:",chardata"`
			W    string `xml:"w,attr"`
			H    string `xml:"h,attr"`
		} `xml:"pgSz"`
		PgMar struct {
			Text   string `xml:",chardata"`
			Top    string `xml:"top,attr"`
			Right  string `xml:"right,attr"`
			Bottom string `xml:"bottom,attr"`
			Left   string `xml:"left,attr"`
			Header string `xml:"header,attr"`
			Footer string `xml:"footer,attr"`
			Gutter string `xml:"gutter,attr"`
		} `xml:"pgMar"`
		Cols struct {
			Text  string `xml:",chardata"`
			Space string `xml:"space,attr"`
		} `xml:"cols"`
		DocGrid struct {
			Text      string `xml:",chardata"`
			LinePitch string `xml:"linePitch,attr"`
		} `xml:"docGrid"`
	} `xml:"sectPr"`
}
