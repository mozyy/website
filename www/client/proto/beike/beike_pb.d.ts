import * as jspb from "google-protobuf"

export class BEIKE extends jspb.Message {
  getCode(): number;
  setCode(value: number): void;

  getData(): BEIKE.DATA | undefined;
  setData(value?: BEIKE.DATA): void;
  hasData(): boolean;
  clearData(): void;

  getMsg(): string;
  setMsg(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BEIKE.AsObject;
  static toObject(includeInstance: boolean, msg: BEIKE): BEIKE.AsObject;
  static serializeBinaryToWriter(message: BEIKE, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BEIKE;
  static deserializeBinaryFromReader(message: BEIKE, reader: jspb.BinaryReader): BEIKE;
}

export namespace BEIKE {
  export type AsObject = {
    code: number,
    data?: BEIKE.DATA.AsObject,
    msg: string,
  }

  export class DATA extends jspb.Message {
    getErrno(): number;
    setErrno(value: number): void;

    getMsg(): string;
    setMsg(value: string): void;

    getData(): BEIKE.DATA.DATA | undefined;
    setData(value?: BEIKE.DATA.DATA): void;
    hasData(): boolean;
    clearData(): void;

    getErrkeysList(): Array<string>;
    setErrkeysList(value: Array<string>): void;
    clearErrkeysList(): void;
    addErrkeys(value: string, index?: number): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DATA.AsObject;
    static toObject(includeInstance: boolean, msg: DATA): DATA.AsObject;
    static serializeBinaryToWriter(message: DATA, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DATA;
    static deserializeBinaryFromReader(message: DATA, reader: jspb.BinaryReader): DATA;
  }

  export namespace DATA {
    export type AsObject = {
      errno: number,
      msg: string,
      data?: BEIKE.DATA.DATA.AsObject,
      errkeysList: Array<string>,
    }

    export class DATA extends jspb.Message {
      getGetershoufanglist(): BEIKE.DATA.DATA.GETERSHOUFANGLIST | undefined;
      setGetershoufanglist(value?: BEIKE.DATA.DATA.GETERSHOUFANGLIST): void;
      hasGetershoufanglist(): boolean;
      clearGetershoufanglist(): void;

      serializeBinary(): Uint8Array;
      toObject(includeInstance?: boolean): DATA.AsObject;
      static toObject(includeInstance: boolean, msg: DATA): DATA.AsObject;
      static serializeBinaryToWriter(message: DATA, writer: jspb.BinaryWriter): void;
      static deserializeBinary(bytes: Uint8Array): DATA;
      static deserializeBinaryFromReader(message: DATA, reader: jspb.BinaryReader): DATA;
    }

    export namespace DATA {
      export type AsObject = {
        getershoufanglist?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.AsObject,
      }

      export class GETERSHOUFANGLIST extends jspb.Message {
        getTotalcount(): number;
        setTotalcount(value: number): void;

        getReturncount(): number;
        setReturncount(value: number): void;

        getFbqueryid(): string;
        setFbqueryid(value: string): void;

        getCurpage(): number;
        setCurpage(value: number): void;

        getHasmoredata(): number;
        setHasmoredata(value: number): void;

        getSpellcheck(): BEIKE.DATA.DATA.GETERSHOUFANGLIST.SPELLCHECK | undefined;
        setSpellcheck(value?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.SPELLCHECK): void;
        hasSpellcheck(): boolean;
        clearSpellcheck(): void;

        getAladincard(): BEIKE.DATA.DATA.GETERSHOUFANGLIST.ALADINCARD | undefined;
        setAladincard(value?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.ALADINCARD): void;
        hasAladincard(): boolean;
        clearAladincard(): void;

        getAdxparams(): BEIKE.DATA.DATA.GETERSHOUFANGLIST.ADXPARAMS | undefined;
        setAdxparams(value?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.ADXPARAMS): void;
        hasAdxparams(): boolean;
        clearAdxparams(): void;

        getNoresultliuzi(): BEIKE.DATA.DATA.GETERSHOUFANGLIST.NORESULTLIUZI | undefined;
        setNoresultliuzi(value?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.NORESULTLIUZI): void;
        hasNoresultliuzi(): boolean;
        clearNoresultliuzi(): void;

        getNoresultreason(): number;
        setNoresultreason(value: number): void;

        getLessresultreason(): number;
        setLessresultreason(value: number): void;

        getListList(): Array<BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST>;
        setListList(value: Array<BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST>): void;
        clearListList(): void;
        addList(value?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST, index?: number): BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST;

        getSelected(): BEIKE.DATA.DATA.GETERSHOUFANGLIST.SELECTED | undefined;
        setSelected(value?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.SELECTED): void;
        hasSelected(): boolean;
        clearSelected(): void;

        getKeyword(): string;
        setKeyword(value: string): void;

        getRecommendextinfo(): BEIKE.DATA.DATA.GETERSHOUFANGLIST.RECOMMENDEXTINFO | undefined;
        setRecommendextinfo(value?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.RECOMMENDEXTINFO): void;
        hasRecommendextinfo(): boolean;
        clearRecommendextinfo(): void;

        getNearwenan(): string;
        setNearwenan(value: string): void;

        getBsrecommendlistList(): Array<string>;
        setBsrecommendlistList(value: Array<string>): void;
        clearBsrecommendlistList(): void;
        addBsrecommendlist(value: string, index?: number): void;

        getTdk(): BEIKE.DATA.DATA.GETERSHOUFANGLIST.TDK | undefined;
        setTdk(value?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.TDK): void;
        hasTdk(): boolean;
        clearTdk(): void;

        getBreadcrumbsList(): Array<BEIKE.DATA.DATA.GETERSHOUFANGLIST.BREADCRUMBS>;
        setBreadcrumbsList(value: Array<BEIKE.DATA.DATA.GETERSHOUFANGLIST.BREADCRUMBS>): void;
        clearBreadcrumbsList(): void;
        addBreadcrumbs(value?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.BREADCRUMBS, index?: number): BEIKE.DATA.DATA.GETERSHOUFANGLIST.BREADCRUMBS;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): GETERSHOUFANGLIST.AsObject;
        static toObject(includeInstance: boolean, msg: GETERSHOUFANGLIST): GETERSHOUFANGLIST.AsObject;
        static serializeBinaryToWriter(message: GETERSHOUFANGLIST, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): GETERSHOUFANGLIST;
        static deserializeBinaryFromReader(message: GETERSHOUFANGLIST, reader: jspb.BinaryReader): GETERSHOUFANGLIST;
      }

      export namespace GETERSHOUFANGLIST {
        export type AsObject = {
          totalcount: number,
          returncount: number,
          fbqueryid: string,
          curpage: number,
          hasmoredata: number,
          spellcheck?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.SPELLCHECK.AsObject,
          aladincard?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.ALADINCARD.AsObject,
          adxparams?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.ADXPARAMS.AsObject,
          noresultliuzi?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.NORESULTLIUZI.AsObject,
          noresultreason: number,
          lessresultreason: number,
          listList: Array<BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST.AsObject>,
          selected?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.SELECTED.AsObject,
          keyword: string,
          recommendextinfo?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.RECOMMENDEXTINFO.AsObject,
          nearwenan: string,
          bsrecommendlistList: Array<string>,
          tdk?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.TDK.AsObject,
          breadcrumbsList: Array<BEIKE.DATA.DATA.GETERSHOUFANGLIST.BREADCRUMBS.AsObject>,
        }

        export class SPELLCHECK extends jspb.Message {
          serializeBinary(): Uint8Array;
          toObject(includeInstance?: boolean): SPELLCHECK.AsObject;
          static toObject(includeInstance: boolean, msg: SPELLCHECK): SPELLCHECK.AsObject;
          static serializeBinaryToWriter(message: SPELLCHECK, writer: jspb.BinaryWriter): void;
          static deserializeBinary(bytes: Uint8Array): SPELLCHECK;
          static deserializeBinaryFromReader(message: SPELLCHECK, reader: jspb.BinaryReader): SPELLCHECK;
        }

        export namespace SPELLCHECK {
          export type AsObject = {
          }
        }


        export class ALADINCARD extends jspb.Message {
          serializeBinary(): Uint8Array;
          toObject(includeInstance?: boolean): ALADINCARD.AsObject;
          static toObject(includeInstance: boolean, msg: ALADINCARD): ALADINCARD.AsObject;
          static serializeBinaryToWriter(message: ALADINCARD, writer: jspb.BinaryWriter): void;
          static deserializeBinary(bytes: Uint8Array): ALADINCARD;
          static deserializeBinaryFromReader(message: ALADINCARD, reader: jspb.BinaryReader): ALADINCARD;
        }

        export namespace ALADINCARD {
          export type AsObject = {
          }
        }


        export class ADXPARAMS extends jspb.Message {
          serializeBinary(): Uint8Array;
          toObject(includeInstance?: boolean): ADXPARAMS.AsObject;
          static toObject(includeInstance: boolean, msg: ADXPARAMS): ADXPARAMS.AsObject;
          static serializeBinaryToWriter(message: ADXPARAMS, writer: jspb.BinaryWriter): void;
          static deserializeBinary(bytes: Uint8Array): ADXPARAMS;
          static deserializeBinaryFromReader(message: ADXPARAMS, reader: jspb.BinaryReader): ADXPARAMS;
        }

        export namespace ADXPARAMS {
          export type AsObject = {
          }
        }


        export class NORESULTLIUZI extends jspb.Message {
          serializeBinary(): Uint8Array;
          toObject(includeInstance?: boolean): NORESULTLIUZI.AsObject;
          static toObject(includeInstance: boolean, msg: NORESULTLIUZI): NORESULTLIUZI.AsObject;
          static serializeBinaryToWriter(message: NORESULTLIUZI, writer: jspb.BinaryWriter): void;
          static deserializeBinary(bytes: Uint8Array): NORESULTLIUZI;
          static deserializeBinaryFromReader(message: NORESULTLIUZI, reader: jspb.BinaryReader): NORESULTLIUZI;
        }

        export namespace NORESULTLIUZI {
          export type AsObject = {
          }
        }


        export class LIST extends jspb.Message {
          getCardtype(): string;
          setCardtype(value: string): void;

          getCityid(): string;
          setCityid(value: string): void;

          getHousecode(): string;
          setHousecode(value: string): void;

          getResblockid(): string;
          setResblockid(value: string): void;

          getDelegateid(): string;
          setDelegateid(value: string): void;

          getTitle(): string;
          setTitle(value: string): void;

          getDesc(): string;
          setDesc(value: string): void;

          getBangdantitle(): string;
          setBangdantitle(value: string): void;

          getRecodesc(): string;
          setRecodesc(value: string): void;

          getTotalprice(): string;
          setTotalprice(value: string): void;

          getTotalpricecount(): string;
          setTotalpricecount(value: string): void;

          getUnitprice(): string;
          setUnitprice(value: string): void;

          getJumpurl(): string;
          setJumpurl(value: string): void;

          getListpictureurl(): BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST.LISTPICTUREURL | undefined;
          setListpictureurl(value?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST.LISTPICTUREURL): void;
          hasListpictureurl(): boolean;
          clearListpictureurl(): void;

          getRecommendreason(): BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST.RECOMMENDREASON | undefined;
          setRecommendreason(value?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST.RECOMMENDREASON): void;
          hasRecommendreason(): boolean;
          clearRecommendreason(): void;

          getColortagsList(): Array<BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST.COLORTAGS>;
          setColortagsList(value: Array<BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST.COLORTAGS>): void;
          clearColortagsList(): void;
          addColortags(value?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST.COLORTAGS, index?: number): BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST.COLORTAGS;

          getHousestatus(): string;
          setHousestatus(value: string): void;

          getIsctypehouse(): boolean;
          setIsctypehouse(value: boolean): void;

          getFbexpoid(): string;
          setFbexpoid(value: string): void;

          getStatusswitch(): BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST.STATUSSWITCH | undefined;
          setStatusswitch(value?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST.STATUSSWITCH): void;
          hasStatusswitch(): boolean;
          clearStatusswitch(): void;

          getBrandtags(): BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST.BRANDTAGS | undefined;
          setBrandtags(value?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST.BRANDTAGS): void;
          hasBrandtags(): boolean;
          clearBrandtags(): void;

          getRecomtagList(): Array<string>;
          setRecomtagList(value: Array<string>): void;
          clearRecomtagList(): void;
          addRecomtag(value: string, index?: number): void;

          getHottop(): BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST.HOTTOP | undefined;
          setHottop(value?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST.HOTTOP): void;
          hasHottop(): boolean;
          clearHottop(): void;

          getFramepictureList(): Array<string>;
          setFramepictureList(value: Array<string>): void;
          clearFramepictureList(): void;
          addFramepicture(value: string, index?: number): void;

          getHousepictureList(): Array<string>;
          setHousepictureList(value: Array<string>): void;
          clearHousepictureList(): void;
          addHousepicture(value: string, index?: number): void;

          getMarketboothList(): Array<string>;
          setMarketboothList(value: Array<string>): void;
          clearMarketboothList(): void;
          addMarketbooth(value: string, index?: number): void;

          getPoidistance(): string;
          setPoidistance(value: string): void;

          getPoiicon(): string;
          setPoiicon(value: string): void;

          serializeBinary(): Uint8Array;
          toObject(includeInstance?: boolean): LIST.AsObject;
          static toObject(includeInstance: boolean, msg: LIST): LIST.AsObject;
          static serializeBinaryToWriter(message: LIST, writer: jspb.BinaryWriter): void;
          static deserializeBinary(bytes: Uint8Array): LIST;
          static deserializeBinaryFromReader(message: LIST, reader: jspb.BinaryReader): LIST;
        }

        export namespace LIST {
          export type AsObject = {
            cardtype: string,
            cityid: string,
            housecode: string,
            resblockid: string,
            delegateid: string,
            title: string,
            desc: string,
            bangdantitle: string,
            recodesc: string,
            totalprice: string,
            totalpricecount: string,
            unitprice: string,
            jumpurl: string,
            listpictureurl?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST.LISTPICTUREURL.AsObject,
            recommendreason?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST.RECOMMENDREASON.AsObject,
            colortagsList: Array<BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST.COLORTAGS.AsObject>,
            housestatus: string,
            isctypehouse: boolean,
            fbexpoid: string,
            statusswitch?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST.STATUSSWITCH.AsObject,
            brandtags?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST.BRANDTAGS.AsObject,
            recomtagList: Array<string>,
            hottop?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.LIST.HOTTOP.AsObject,
            framepictureList: Array<string>,
            housepictureList: Array<string>,
            marketboothList: Array<string>,
            poidistance: string,
            poiicon: string,
          }

          export class LISTPICTUREURL extends jspb.Message {
            serializeBinary(): Uint8Array;
            toObject(includeInstance?: boolean): LISTPICTUREURL.AsObject;
            static toObject(includeInstance: boolean, msg: LISTPICTUREURL): LISTPICTUREURL.AsObject;
            static serializeBinaryToWriter(message: LISTPICTUREURL, writer: jspb.BinaryWriter): void;
            static deserializeBinary(bytes: Uint8Array): LISTPICTUREURL;
            static deserializeBinaryFromReader(message: LISTPICTUREURL, reader: jspb.BinaryReader): LISTPICTUREURL;
          }

          export namespace LISTPICTUREURL {
            export type AsObject = {
            }
          }


          export class RECOMMENDREASON extends jspb.Message {
            serializeBinary(): Uint8Array;
            toObject(includeInstance?: boolean): RECOMMENDREASON.AsObject;
            static toObject(includeInstance: boolean, msg: RECOMMENDREASON): RECOMMENDREASON.AsObject;
            static serializeBinaryToWriter(message: RECOMMENDREASON, writer: jspb.BinaryWriter): void;
            static deserializeBinary(bytes: Uint8Array): RECOMMENDREASON;
            static deserializeBinaryFromReader(message: RECOMMENDREASON, reader: jspb.BinaryReader): RECOMMENDREASON;
          }

          export namespace RECOMMENDREASON {
            export type AsObject = {
            }
          }


          export class COLORTAGS extends jspb.Message {
            getKey(): string;
            setKey(value: string): void;

            getTitle(): string;
            setTitle(value: string): void;

            getColor(): string;
            setColor(value: string): void;

            getTextcolor(): string;
            setTextcolor(value: string): void;

            getBgcolor(): string;
            setBgcolor(value: string): void;

            serializeBinary(): Uint8Array;
            toObject(includeInstance?: boolean): COLORTAGS.AsObject;
            static toObject(includeInstance: boolean, msg: COLORTAGS): COLORTAGS.AsObject;
            static serializeBinaryToWriter(message: COLORTAGS, writer: jspb.BinaryWriter): void;
            static deserializeBinary(bytes: Uint8Array): COLORTAGS;
            static deserializeBinaryFromReader(message: COLORTAGS, reader: jspb.BinaryReader): COLORTAGS;
          }

          export namespace COLORTAGS {
            export type AsObject = {
              key: string,
              title: string,
              color: string,
              textcolor: string,
              bgcolor: string,
            }
          }


          export class STATUSSWITCH extends jspb.Message {
            getIsyezhutuijian(): boolean;
            setIsyezhutuijian(value: boolean): void;

            getIshaofang(): boolean;
            setIshaofang(value: boolean): void;

            getIsyezhupay(): boolean;
            setIsyezhupay(value: boolean): void;

            getIsvr(): boolean;
            setIsvr(value: boolean): void;

            getIskey(): boolean;
            setIskey(value: boolean): void;

            getIsnew(): boolean;
            setIsnew(value: boolean): void;

            serializeBinary(): Uint8Array;
            toObject(includeInstance?: boolean): STATUSSWITCH.AsObject;
            static toObject(includeInstance: boolean, msg: STATUSSWITCH): STATUSSWITCH.AsObject;
            static serializeBinaryToWriter(message: STATUSSWITCH, writer: jspb.BinaryWriter): void;
            static deserializeBinary(bytes: Uint8Array): STATUSSWITCH;
            static deserializeBinaryFromReader(message: STATUSSWITCH, reader: jspb.BinaryReader): STATUSSWITCH;
          }

          export namespace STATUSSWITCH {
            export type AsObject = {
              isyezhutuijian: boolean,
              ishaofang: boolean,
              isyezhupay: boolean,
              isvr: boolean,
              iskey: boolean,
              isnew: boolean,
            }
          }


          export class BRANDTAGS extends jspb.Message {
            serializeBinary(): Uint8Array;
            toObject(includeInstance?: boolean): BRANDTAGS.AsObject;
            static toObject(includeInstance: boolean, msg: BRANDTAGS): BRANDTAGS.AsObject;
            static serializeBinaryToWriter(message: BRANDTAGS, writer: jspb.BinaryWriter): void;
            static deserializeBinary(bytes: Uint8Array): BRANDTAGS;
            static deserializeBinaryFromReader(message: BRANDTAGS, reader: jspb.BinaryReader): BRANDTAGS;
          }

          export namespace BRANDTAGS {
            export type AsObject = {
            }
          }


          export class HOTTOP extends jspb.Message {
            getDspagentucid(): string;
            setDspagentucid(value: string): void;

            getHottopdigv(): string;
            setHottopdigv(value: string): void;

            getHottop(): number;
            setHottop(value: number): void;

            serializeBinary(): Uint8Array;
            toObject(includeInstance?: boolean): HOTTOP.AsObject;
            static toObject(includeInstance: boolean, msg: HOTTOP): HOTTOP.AsObject;
            static serializeBinaryToWriter(message: HOTTOP, writer: jspb.BinaryWriter): void;
            static deserializeBinary(bytes: Uint8Array): HOTTOP;
            static deserializeBinaryFromReader(message: HOTTOP, reader: jspb.BinaryReader): HOTTOP;
          }

          export namespace HOTTOP {
            export type AsObject = {
              dspagentucid: string,
              hottopdigv: string,
              hottop: number,
            }
          }

        }


        export class SELECTED extends jspb.Message {
          getRegion(): BEIKE.DATA.DATA.GETERSHOUFANGLIST.SELECTED.REGION | undefined;
          setRegion(value?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.SELECTED.REGION): void;
          hasRegion(): boolean;
          clearRegion(): void;

          getPrice(): BEIKE.DATA.DATA.GETERSHOUFANGLIST.SELECTED.PRICE | undefined;
          setPrice(value?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.SELECTED.PRICE): void;
          hasPrice(): boolean;
          clearPrice(): void;

          getRoom(): BEIKE.DATA.DATA.GETERSHOUFANGLIST.SELECTED.ROOM | undefined;
          setRoom(value?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.SELECTED.ROOM): void;
          hasRoom(): boolean;
          clearRoom(): void;

          getMore(): BEIKE.DATA.DATA.GETERSHOUFANGLIST.SELECTED.MORE | undefined;
          setMore(value?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.SELECTED.MORE): void;
          hasMore(): boolean;
          clearMore(): void;

          getOrder(): BEIKE.DATA.DATA.GETERSHOUFANGLIST.SELECTED.ORDER | undefined;
          setOrder(value?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.SELECTED.ORDER): void;
          hasOrder(): boolean;
          clearOrder(): void;

          serializeBinary(): Uint8Array;
          toObject(includeInstance?: boolean): SELECTED.AsObject;
          static toObject(includeInstance: boolean, msg: SELECTED): SELECTED.AsObject;
          static serializeBinaryToWriter(message: SELECTED, writer: jspb.BinaryWriter): void;
          static deserializeBinary(bytes: Uint8Array): SELECTED;
          static deserializeBinaryFromReader(message: SELECTED, reader: jspb.BinaryReader): SELECTED;
        }

        export namespace SELECTED {
          export type AsObject = {
            region?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.SELECTED.REGION.AsObject,
            price?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.SELECTED.PRICE.AsObject,
            room?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.SELECTED.ROOM.AsObject,
            more?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.SELECTED.MORE.AsObject,
            order?: BEIKE.DATA.DATA.GETERSHOUFANGLIST.SELECTED.ORDER.AsObject,
          }

          export class REGION extends jspb.Message {
            getIdList(): Array<string>;
            setIdList(value: Array<string>): void;
            clearIdList(): void;
            addId(value: string, index?: number): void;

            getName(): string;
            setName(value: string): void;

            getLev3(): string;
            setLev3(value: string): void;

            getLev1(): string;
            setLev1(value: string): void;

            getLev2(): string;
            setLev2(value: string): void;

            serializeBinary(): Uint8Array;
            toObject(includeInstance?: boolean): REGION.AsObject;
            static toObject(includeInstance: boolean, msg: REGION): REGION.AsObject;
            static serializeBinaryToWriter(message: REGION, writer: jspb.BinaryWriter): void;
            static deserializeBinary(bytes: Uint8Array): REGION;
            static deserializeBinaryFromReader(message: REGION, reader: jspb.BinaryReader): REGION;
          }

          export namespace REGION {
            export type AsObject = {
              idList: Array<string>,
              name: string,
              lev3: string,
              lev1: string,
              lev2: string,
            }
          }


          export class PRICE extends jspb.Message {
            serializeBinary(): Uint8Array;
            toObject(includeInstance?: boolean): PRICE.AsObject;
            static toObject(includeInstance: boolean, msg: PRICE): PRICE.AsObject;
            static serializeBinaryToWriter(message: PRICE, writer: jspb.BinaryWriter): void;
            static deserializeBinary(bytes: Uint8Array): PRICE;
            static deserializeBinaryFromReader(message: PRICE, reader: jspb.BinaryReader): PRICE;
          }

          export namespace PRICE {
            export type AsObject = {
            }
          }


          export class ROOM extends jspb.Message {
            serializeBinary(): Uint8Array;
            toObject(includeInstance?: boolean): ROOM.AsObject;
            static toObject(includeInstance: boolean, msg: ROOM): ROOM.AsObject;
            static serializeBinaryToWriter(message: ROOM, writer: jspb.BinaryWriter): void;
            static deserializeBinary(bytes: Uint8Array): ROOM;
            static deserializeBinaryFromReader(message: ROOM, reader: jspb.BinaryReader): ROOM;
          }

          export namespace ROOM {
            export type AsObject = {
            }
          }


          export class MORE extends jspb.Message {
            serializeBinary(): Uint8Array;
            toObject(includeInstance?: boolean): MORE.AsObject;
            static toObject(includeInstance: boolean, msg: MORE): MORE.AsObject;
            static serializeBinaryToWriter(message: MORE, writer: jspb.BinaryWriter): void;
            static deserializeBinary(bytes: Uint8Array): MORE;
            static deserializeBinaryFromReader(message: MORE, reader: jspb.BinaryReader): MORE;
          }

          export namespace MORE {
            export type AsObject = {
            }
          }


          export class ORDER extends jspb.Message {
            serializeBinary(): Uint8Array;
            toObject(includeInstance?: boolean): ORDER.AsObject;
            static toObject(includeInstance: boolean, msg: ORDER): ORDER.AsObject;
            static serializeBinaryToWriter(message: ORDER, writer: jspb.BinaryWriter): void;
            static deserializeBinary(bytes: Uint8Array): ORDER;
            static deserializeBinaryFromReader(message: ORDER, reader: jspb.BinaryReader): ORDER;
          }

          export namespace ORDER {
            export type AsObject = {
            }
          }

        }


        export class RECOMMENDEXTINFO extends jspb.Message {
          getBizcircleidList(): Array<string>;
          setBizcircleidList(value: Array<string>): void;
          clearBizcircleidList(): void;
          addBizcircleid(value: string, index?: number): void;

          serializeBinary(): Uint8Array;
          toObject(includeInstance?: boolean): RECOMMENDEXTINFO.AsObject;
          static toObject(includeInstance: boolean, msg: RECOMMENDEXTINFO): RECOMMENDEXTINFO.AsObject;
          static serializeBinaryToWriter(message: RECOMMENDEXTINFO, writer: jspb.BinaryWriter): void;
          static deserializeBinary(bytes: Uint8Array): RECOMMENDEXTINFO;
          static deserializeBinaryFromReader(message: RECOMMENDEXTINFO, reader: jspb.BinaryReader): RECOMMENDEXTINFO;
        }

        export namespace RECOMMENDEXTINFO {
          export type AsObject = {
            bizcircleidList: Array<string>,
          }
        }


        export class TDK extends jspb.Message {
          getTitle(): string;
          setTitle(value: string): void;

          getKeywords(): string;
          setKeywords(value: string): void;

          getDescription(): string;
          setDescription(value: string): void;

          serializeBinary(): Uint8Array;
          toObject(includeInstance?: boolean): TDK.AsObject;
          static toObject(includeInstance: boolean, msg: TDK): TDK.AsObject;
          static serializeBinaryToWriter(message: TDK, writer: jspb.BinaryWriter): void;
          static deserializeBinary(bytes: Uint8Array): TDK;
          static deserializeBinaryFromReader(message: TDK, reader: jspb.BinaryReader): TDK;
        }

        export namespace TDK {
          export type AsObject = {
            title: string,
            keywords: string,
            description: string,
          }
        }


        export class BREADCRUMBS extends jspb.Message {
          getTitle(): string;
          setTitle(value: string): void;

          getUrl(): string;
          setUrl(value: string): void;

          serializeBinary(): Uint8Array;
          toObject(includeInstance?: boolean): BREADCRUMBS.AsObject;
          static toObject(includeInstance: boolean, msg: BREADCRUMBS): BREADCRUMBS.AsObject;
          static serializeBinaryToWriter(message: BREADCRUMBS, writer: jspb.BinaryWriter): void;
          static deserializeBinary(bytes: Uint8Array): BREADCRUMBS;
          static deserializeBinaryFromReader(message: BREADCRUMBS, reader: jspb.BinaryReader): BREADCRUMBS;
        }

        export namespace BREADCRUMBS {
          export type AsObject = {
            title: string,
            url: string,
          }
        }

      }

    }

  }

}

